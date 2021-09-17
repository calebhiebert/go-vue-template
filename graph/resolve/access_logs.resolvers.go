package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models"
)

func (r *accessLogResolver) RequestHeaders(ctx context.Context, obj *models.AccessLog) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accessLogResolver) ResponseBody(ctx context.Context, obj *models.AccessLog) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accessLogResolver) ResponseHeaders(ctx context.Context, obj *models.AccessLog) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// AccessLog returns generated.AccessLogResolver implementation.
func (r *Resolver) AccessLog() generated.AccessLogResolver { return &accessLogResolver{r} }

type accessLogResolver struct{ *Resolver }
