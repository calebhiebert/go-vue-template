package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models"
)

func (r *courseResolver) Latitude(ctx context.Context, obj *models.Course) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *courseResolver) Longitude(ctx context.Context, obj *models.Course) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

type courseResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *courseResolver) Lat(ctx context.Context, obj *models.Course) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *courseResolver) Lng(ctx context.Context, obj *models.Course) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}
