package docs

// swagger:model
type Workout struct {
	// The unique integer ID of an athletes workout
	// Read Only: true
	Workout_ID  int    `json:"workout_id"`
	// The unique ID that links an athlete to a workout
	// Read Only: true
	Athlete_ID  string `json:"-" validate:"required"`
	// The date the workout was completed
	// required: true
	Date        string `json:"date"`
	// The description of the workout, how the athlete felt, what they did, etc
	// required: true
	Description string `json:"description" validate:"required"`
	// The sport/excercise the athlete did
	// required: true
	Sport       string `json:"sport"`
	// The athletes rating of the workout
	// min: 0
	// max: 10
	// required: true
	Rating      int    `json:"rating"`
}

// swagger:model
type Athlete struct {
	// The unique ID that distinguishes athletes
	// Read Only: true
	Athlete_ID string    `json:"athlete_id"`
	// The first name of the athlete
	// required: true
	First_Name string    `json:"first_name" validate:"required"`
	// The last name of the athlete
	// required: true
	Last_Name  string    `json:"last_name" validate:"required"`
	// The age of the athlete
	// min: 18
	// required: true
	Age        uint8     `json:"age"`
	// The date the athlete joined AER
	// Read Only: true
	Joined     string `json:"joined"`
}

// swagger:model
type Efficiency struct{
	// The unique ID that distinguishes efficiency objects from each other
	// Read Only: true
	Efficiency_ID    int     `json:"-"`
	// The efficiency score of an athlete based upon their past workouts
	// Example: 20.34
	// Read Only: true
	Efficiency_Score float64 `json:"efficiency_score"`
	// The favorite sport of the athlete based upon what they do in workouts most often
	// Example: swimming
	// Read Only: true
	Favorite_Sport   string  `json:"favorite_excercise"`
	// The link between the athlete and efficiency score
	// Read Only: true
	Athlete_ID       string  `json:"-"`
}