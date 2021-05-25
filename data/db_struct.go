package data

import(
	"database/sql"
	"fmt"
	"log"
	"strings"
	_ "github.com/lib/pq"
)

type Efficiency struct {
	Efficiency_ID    int     `json:"efficiency_id"`
	Efficiency_Score float64 `json:"efficiency_score"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Athlete_ID       string  `json:"athlete_id"`
}

//Workout Structure which holds details about an athletes workout
type Workout struct {
	Workout_ID  int    `json:"workout_id"`
	Athlete_ID  string `json:"athlete_id" validate:"required"`
	Date        string `json:"date"`
	Description string `json:"description" validate:"required"`
	Sport       string `json:"sport" validate isListed`
	Rating      int    `json:"rating" validate: gte=0, lte=10`
}

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	Athlete_ID int    `json:"athlete_id"`
	First_Name string `json:"first_name" validate:"required"`
	Last_Name  string `json:"last_name" validate:"required"`
	Age        uint8  `json:"age" validate:"gte=18, lte=100`
	Joined     string `json:"joined"`
}


type EfficiencyViewWrapper struct {
	First_Name string `json:"first_name" validate:"required"`
	Last_Name  string `json:"last_name" validate:"required"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Efficiency_Score float64 `json:"efficiency_score"`
}
