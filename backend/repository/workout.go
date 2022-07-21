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
	if err := wr.db.WithContext(ctx).Find(&workouts).Error; err != nil {
		return nil, err
	}
	for _, w := range workouts {
		excercises, err := wr.getExcercises(ctx, w.WorkoutID)
		if err != nil {
			return nil, err
		}
		for _, e := range excercises {
			sets, err := wr.getSets(ctx, e.ID)
			if err != nil {
				return nil, err
			}
			e.Sets = sets
		}
		w.Excercises = excercises
	}
	return workouts, nil
}

func (wr *workoutRepository) getExcercises(ctx context.Context, wid int) ([]*models.Excercise, error) {
	var excercises []*models.Excercise
	if err := wr.db.WithContext(ctx).Find(excercises).Error; err != nil {
		return nil, err
	}
	return excercises, nil
}

func (wr *workoutRepository) getSets(ctx context.Context, excID int) ([]*models.Sets, error) {
	var sets []*models.Sets
	if err := wr.db.WithContext(ctx).Find(sets).Error; err != nil {
		return nil, err
	}
	return sets, nil
}

func (wr *workoutRepository) GetByID(ctx context.Context, wid int) (*models.Workout, error) {
	var workout *models.Workout
	if err := wr.db.WithContext(ctx).Where("id = @id", sql.Named("id", wid)).Find(workout).Error; err != nil {
		return nil, err
	}
	return workout, nil
}
