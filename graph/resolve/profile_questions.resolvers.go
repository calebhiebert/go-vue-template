package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models"
)

func (r *profileQuestionResolver) Options(ctx context.Context, obj *models.ProfileQuestion) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// ProfileQuestion returns generated.ProfileQuestionResolver implementation.
func (r *Resolver) ProfileQuestion() generated.ProfileQuestionResolver {
	return &profileQuestionResolver{r}
}

type profileQuestionResolver struct{ *Resolver }
