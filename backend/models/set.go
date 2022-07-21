package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Sets struct {
	Id     int     `json:"id" gorm:"primaryKey"`
	Weight float32 `json:"weight" gorm:"type:int; not null"`
	Reps   int     `json:"reps" gorm:"type:int; not null"`
}

func (s Sets) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Reps, validation.Required),
		validation.Field(&s.Weight, validation.Required),
	)
}
