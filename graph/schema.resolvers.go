package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/slowmanchan/grapqlGunpla/graph/generated"
	"github.com/slowmanchan/grapqlGunpla/graph/model"
	"github.com/slowmanchan/grapqlGunpla/internal/gunplas"
)

func (r *queryResolver) Gunplas(ctx context.Context) ([]*model.Gunpla, error) {
	var resultGunplas []*model.Gunpla
	var dbGunplas []gunplas.Gunpla
	dbGunplas = gunplas.GetAll()

	for _, gunpla := range dbGunplas {
		resultGunplas = append(resultGunplas, &model.Gunpla{
			ID:              gunpla.ID,
			BoxArtImageLink: gunpla.BoxArtImageLink,
			Title:           gunpla.Title,
			Subtitle:        gunpla.Subtitle,
			Classification:  gunpla.Classification,
			LineupNo:        gunpla.LineupNo,
			Scale:           gunpla.Scale,
			Franchise:       gunpla.Franchise,
			ReleaseDate:     gunpla.ReleaseDate,
			JanIsbn:         gunpla.JanIsbn,
			Run:             gunpla.Run,
			Price:           gunpla.Price,
			Includes:        gunpla.Includes,
			Features:        gunpla.Franchise,
		})
	}
	return resultGunplas, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
