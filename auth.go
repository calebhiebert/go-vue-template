package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/cristalhq/jwt"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

const BcryptCost = 14

// TODO obtain key more safely
const JWTKey = "supersecret"

type AuthenticationResult struct {
	JWT  string      `json:"jwt"`
	User models.User `json:"user"`
}

type JWTClaims struct {
	jwt.StandardClaims
	User models.User `json:"user"`
}

func (c *JWTClaims) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,max=30"`
	Login    string `json:"login" binding:"required,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=128"`
}

// RegisterUsernamePassword godoc
// @Summary Registers a new account using username and password as the authentication scheme
// @Produce json
// @Accept json
// @Param req body RegisterRequest true "Register parameters"
// @Success 200 {object} AuthenticationResult
// @Router /auth/register [post]
func (*Controller) RegisterUsernamePassword(c *gin.Context) {
	var req RegisterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	// Check for existing users with the same login
	loginOrEmailExists, err := models.Users(qm.Where("login = ?", req.Login), qm.Or("email = ?", req.Email)).
		ExistsG(c.Request.Context())
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	if loginOrEmailExists {
		NewAPIError("email-or-login-exists", http.StatusBadRequest, "A user with that email or login already exists").
			Respond(c)
		return
	}

	// Hash user's desired password
	pwHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), BcryptCost)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	newUser := models.User{
		ID:     uuid.Must(uuid.NewV4()).String(),
		Name:   req.Name,
		Login:  null.StringFrom(req.Login),
		Email:  req.Email,
		PWHash: null.StringFrom(string(pwHash)),
	}

	// Insert the user into the database
	err = newUser.InsertG(c.Request.Context(), boil.Infer())
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	generatedToken, claims, err := createJWTForUser(&newUser)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	tokenIssuance := models.TokenIssuance{
		ID:        claims.ID,
		IPAddress: c.ClientIP(),
		UserID:    claims.User.ID,
	}

	err = tokenIssuance.InsertG(c.Request.Context(), boil.Infer())
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, AuthenticationResult{
		JWT:  string(generatedToken.Raw()),
		User: newUser,
	})
}

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthenticateUsernamePassword godoc
// @Summary Exchanges a username and password for a signed JWT
// @Produce json
// @Accept json
// @Param req body LoginRequest true "Login parameters"
// @Success 200 {object} AuthenticationResult
// @Router /auth/loginup [post]
func (*Controller) AuthenticateUsernamePassword(c *gin.Context) {
	var req LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	// Load user from the database
	user, err := models.Users(qm.Where("login = ?", req.Login)).OneG(c.Request.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			NewAPIError("username-or-password-incorrect", http.StatusUnauthorized, "The username or password you entered is incorrect").
				Respond(c)
			return
		}

		APIErrorFromErr(err).Respond(c)
		return
	}

	if user.PWHash.IsZero() {
		NewAPIError("invalid-login-type", http.StatusBadRequest, "This user does not support username and password login").
			Respond(c)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PWHash.String), []byte(req.Password))
	if err != nil {
		NewAPIError("username-or-password-incorrect", http.StatusUnauthorized, "The username or password you entered is incorrect").
			Respond(c)
		return
	}

	generatedToken, claims, err := createJWTForUser(user)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	tokenIssuance := models.TokenIssuance{
		ID:        claims.ID,
		IPAddress: c.ClientIP(),
		UserID:    claims.User.ID,
	}

	err = tokenIssuance.InsertG(c.Request.Context(), boil.Infer())
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, AuthenticationResult{
		JWT:  string(generatedToken.Raw()),
		User: *user,
	})
}

// AuthenticateJWT godoc
// @Summary Exchanges a JWT from a configurable source for a signed JWT
// @Produce json
// @Success 200 {object} AuthenticationResult
// @Router /auth/loginjwt [post]
func (*Controller) AuthenticateJWT(c *gin.Context) {

}

func getJWTSigner() jwt.Signer {
	signer, err := jwt.NewHS256([]byte(JWTKey))
	if err != nil {
		panic(err)
	}

	return signer
}

func createJWTForUser(user *models.User) (*jwt.Token, *JWTClaims, error) {
	builder := jwt.NewTokenBuilder(getJWTSigner())

	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  []string{os.Getenv("HOSTED_URL")},
			ExpiresAt: jwt.Timestamp(time.Now().UTC().Add(48 * time.Hour).Unix()),
			ID:        uuid.Must(uuid.NewV4()).String(),
			IssuedAt:  jwt.Timestamp(time.Now().UTC().Unix()),
			Issuer:    os.Getenv("HOSTED_URL"),
			Subject:   user.ID,
		},
		User: *user,
	}

	token, err := builder.Build(&claims)
	if err != nil {
		return nil, nil, err
	}

	return token, &claims, nil
}

func verifyTokenMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.Next()
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		NewAPIError("invalid-token", http.StatusUnauthorized, "An invalid bearer token was submitted").
			Abort(c)
		return
	}

	parsedToken, err := jwt.ParseAndVerifyString(token, getJWTSigner())
	if err != nil {
		NewAPIError("invalid-token", http.StatusUnauthorized, "Failed to verify token").
			Abort(c)
		return
	}

	var claims JWTClaims

	err = json.Unmarshal(parsedToken.RawClaims(), &claims)
	if err != nil {
		APIErrorFromErr(err).Abort(c)
		return
	}

	if !time.Now().UTC().Before(claims.ExpiresAt.Time()) {
		NewAPIError("token-expired", http.StatusUnauthorized, "The token provided has expired").Abort(c)
		return
	}

	c.Set("verified_user", &claims.User)

	ctxWithUser := context.WithValue(c.Request.Context(), "user", &claims.User)

	c.Request = c.Request.WithContext(ctxWithUser)

	c.Next()
}

func mustBeAuthenticatedMiddleware(c *gin.Context) {
	user := extractVerifiedUser(c)

	if user == nil {
		NewAPIError("must-be-authenticated", http.StatusUnauthorized, "You must be authenticated to access this resource").
			Abort(c)
		return
	}

	c.Next()
}

func userHasRoleMiddleware(role string) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := extractVerifiedUser(c)

		if user == nil {
			NewAPIError("missing-user", http.StatusUnauthorized, "Must be logged in").Abort(c)
			return
		}

		for _, r := range user.Roles {
			if role == r {
				c.Next()
				return
			}
		}

		NewAPIError("missing-user-role", http.StatusForbidden, "User is missing required role " + role).Abort(c)
	}
}

func userHasRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (res interface{}, err error) {
	if user := extractUser(ctx); user == nil || !stringSliceContains(user.Roles, role) {
		return nil, errors.New("User was missing required role")
	}

	return next(ctx)
}

func extractVerifiedUser(c *gin.Context) *models.User {
	user, exists := c.Get("verified_user")
	if !exists {
		return nil
	}

	return user.(*models.User)
}

func extractUser(ctx context.Context) *models.User {
	raw, ok := ctx.Value("user").(*models.User)
	if !ok {
		return nil
	}

	return raw
}
