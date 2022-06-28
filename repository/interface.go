package repository

import (
	"context"

	"github.com/BrandonReno/A.E.R/models"
)

type (
	AthleteRepository interface{
		AddAthlete(ctx context.Context, a *models.Athlete) error
		GetAllAthletes(ctx context.Context) ([]*models.Athlete, error)
		GetAthleteByID(ctx context.Context, aid int) (*models.Athlete, error)
		UpdateAthlete(ctx context.Context, a *models.Athlete) error
	}

	WorkoutRepository interface{
		Create(ctx context.Context, w *models.Workout) error
		GetAll(ctx context.Context) ([]*models.Workout, error)
		GetAllByAthleteID(ctx context.Context, aid int) ([]*models.Workout, error)
		GetByID(ctx context.Context, wid int) (*models.Workout, error)
	}
)