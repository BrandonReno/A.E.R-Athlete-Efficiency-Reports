package repository

import (
	"context"

	"github.com/BrandonReno/A.E.R/models"
)

type (
	AthleteRepository interface {
		Create(ctx context.Context, a *models.Athlete) error
		GetAll(ctx context.Context) ([]*models.Athlete, error)
		GetByID(ctx context.Context, aid int) (*models.Athlete, error)
		Update(ctx context.Context, a *models.Athlete) (*models.Athlete, error)
	}

	WorkoutRepository interface {
		Create(ctx context.Context, w *models.Workout) error
		GetAll(ctx context.Context) ([]*models.Workout, error)
		GetAllByAthleteID(ctx context.Context, aid int) ([]*models.Workout, error)
		GetByID(ctx context.Context, wid int) (*models.Workout, error)
	}
)
