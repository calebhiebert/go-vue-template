package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models"
)

func (r *eventPQMatchScoreResolver) NumAvg(ctx context.Context, obj *models.EventPQMatchScore) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// EventPQMatchScore returns generated.EventPQMatchScoreResolver implementation.
func (r *Resolver) EventPQMatchScore() generated.EventPQMatchScoreResolver {
	return &eventPQMatchScoreResolver{r}
}

type eventPQMatchScoreResolver struct{ *Resolver }
