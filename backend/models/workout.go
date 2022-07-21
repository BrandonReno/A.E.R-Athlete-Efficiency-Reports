package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Workout struct {
	WorkoutID   int          `json:"id,omitempty" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"type:text; not null"`
	Description string       `json:"description" gorm:"type:text; not null"`
	CreatedAt   time.Time    `json:"created_at,omitempty" gorm:"not null;default:now()"`
	Excercises  []*Excercise `json:"excercises" gorm:"-"`
}

func (w Workout) Validate() error{
	return validation.ValidateStruct(&w,
		validation.Field(&w.Title, validation.Required, validation.Length(3, 40)),
		validation.Field(&w.Description, validation.Required, validation.Length(3,200)),
	)
}
