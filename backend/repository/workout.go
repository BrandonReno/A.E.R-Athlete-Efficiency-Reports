package repository

import (
	"context"
	"database/sql"

	"github.com/BrandonReno/A.E.R/models"
	"gorm.io/gorm"
)

type workoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) WorkoutRepository {
	return &workoutRepository{db: db}
}

func (wr *workoutRepository) Create(ctx context.Context, w *models.Workout) error {
	if err := wr.db.WithContext(ctx).Create(w).Error; err != nil {
		return err
	}
	return nil
}

func (wr *workoutRepository) GetAll(ctx context.Context) ([]*models.Workout, error) {
	var workouts []*models.Workout
	if err := wr.db.WithContext(ctx).Find(workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (wr *workoutRepository) GetAllByAthleteID(ctx context.Context, aid int) ([]*models.Workout, error) {
	var workouts []*models.Workout
	if err := wr.db.WithContext(ctx).Where("athlete_id = @athlete_id", sql.Named("athlete_id", aid)).Find(workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (wr *workoutRepository) GetByID(ctx context.Context, wid int) (*models.Workout, error) {
	var workout *models.Workout
	if err := wr.db.WithContext(ctx).Where("id = @id", sql.Named("id", wid)).Find(workout).Error; err != nil {
		return nil, err
	}
	return workout, nil
}
