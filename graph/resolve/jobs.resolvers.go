package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models"
)

func (r *jobResolver) Data(ctx context.Context, obj *models.Job) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *jobResolver) FailureData(ctx context.Context, obj *models.Job) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// Job returns generated.JobResolver implementation.
func (r *Resolver) Job() generated.JobResolver { return &jobResolver{r} }

type jobResolver struct{ *Resolver }
