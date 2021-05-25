package docs

// swagger:model
type Workout struct { 
	// The id for this workout
    //
    // min: 1
	// read only: true
	ID          int     `json:"id"`
	// The user whom completed this workout
    //
    // required: true
	User        Athlete `json:"user" validate:"required"`
	// The date this workout was completed
	//
	// read only: true
	Date        string  `json:"date"`
	// The description of the workout
    //
    // maximum length: 500
	// required: true
	Description string  `json:"description" validate:"required"`
}

//Athlete Structure which defines the 'User' of a workout
type Athlete struct {
	// The ID linking to the user
    //
    // min: 1
	// Read Only: true
	ID    int    `json:"athlete_id"`
	// The full name of the user
    //
    // minimum length: 5
	// maximum length: 50
	// required: true
	Name  string `json:"name" validate:"required"`
	// The sport the user does
    //
	// required: true
	Sport string `json:"sport" validate:"required,Sport"`
	// The age of the user
    //
    // min: 18
	// max: 100
	Age   uint8    `json:"age" validate:"gte=18, lte=100`
}