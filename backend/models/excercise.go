package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Excercise struct {
	ID          int     `json:"id,omitempty" gorm:"primaryKey"`
	Title       string  `json:"title" gorm:"type:text; not null"`
	Description string  `json:"description" gorm:"type:text; not null"`
	WorkoutID   int     `json:"workout_id,omitempty" gorm:"type:integer references workouts(id);not null;unique"`
	Sets        []*Sets `json:"sets" gorm:"-"`
}

func (e Excercise) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required, validation.Length(3, 40)),
		validation.Field(&e.Description, validation.Required, validation.Length(3, 200)),
	)
}
