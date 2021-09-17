package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/calebhiebert/go-vue-template/db"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const UserBatchSize = 50

func main() {
	// Load the .env file (if present)
	err := godotenv.Load("../../.env", ".env")
	if err != nil {
		fmt.Println("Failed to load env file. Most of time this error can be ignored")
	}

	// Create a connection to the postgres database
	dbConn, err := db.SetupDatabase()
	if err != nil {
		panic(err)
	}

	// Setup the global db connection for sqlboiler
	boil.SetDB(dbConn)

	// Generate some users
	// All users will have the password "password"
	SeedUsers(context.Background(), 1000)
}

func SeedUsers(ctx context.Context, userCount int) {
	pw, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	pwHash := string(pw)

	for i := 0; i < userCount; i++ {
		eml := gofakeit.Email()

		user := models.User{
			ID:     uuid.Must(uuid.NewV4()).String(),
			Name:   fmt.Sprintf("%s %s", gofakeit.FirstName(), gofakeit.LastName()),
			Login:  null.StringFrom(eml),
			Email:  eml,
			PWHash: null.StringFrom(pwHash),
		}

		// 75% chance of generating a gender
		if intR(0, 100) <= 75 {
			g := strings.ToLower(gofakeit.Gender())
			gsd := false

			if g != "male" && g != "female" {
				gsd = true
			}

			user.Gender = null.StringFrom(g)
			user.GenderSelfDefined = null.BoolFrom(gsd)
		}

		// 75% chance of generating a birthday
		if intR(0, 100) <= 75 {
			bdayFloor := time.Now().Add(-(80 * 365 * 24) * time.Hour)
			bdayCeil := time.Now().Add(-(18 * 365 * 24) * time.Hour)

			user.Birthday = null.TimeFrom(gofakeit.DateRange(bdayFloor, bdayCeil))
		}

		// 75% chance of generating a phone number
		if intR(0, 100) <= 75 {
			user.Phone = null.StringFrom(gofakeit.Phone())
		}

		err = user.InsertG(ctx, boil.Infer())
		if err != nil {
			panic(err)
		}

		fmt.Println("USER ", i+1, "/", userCount)
	}
}

func intR(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
