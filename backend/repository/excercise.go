package repository 


import (
	"context"
	"database/sql"

	"github.com/BrandonReno/A.E.R/models"
	"gorm.io/gorm"
)

type excerciseRepository struct {
	db *gorm.DB
}

func NewExcerciseRepository(db *gorm.DB) WorkoutRepository {
	return &workoutRepository{db: db}
}