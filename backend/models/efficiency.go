package models

type Efficiency struct {
	Efficiency_ID    int     `json:"-"`
	Efficiency_Score float64 `json:"efficiency_score"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Athlete_ID       string  `json:"-"`
}

func (e *Efficiency) CalculateEfficiency(w []Workout) {
	averageRating := 0

	for _, workout := range w {
		averageRating += workout.Rating
	}

	averageRating = averageRating / len(w)
	lastInd := len(w) - 1

	if w[lastInd].Rating < averageRating {
		e.Efficiency_Score /= float64(w[lastInd].Rating)
	} else {
		e.Efficiency_Score *= float64(w[lastInd].Rating)
	}
}

func (e *Efficiency) CalculateSport(w []Workout) {
	sportMap := make(map[string]int)

	for _, workout := range w {
		sportMap[workout.Sport] += 1
	}

	favoriteSport := ""
	maximum := 0
	for key, val := range sportMap {
		if val > maximum {
			favoriteSport = key
		}
	}

	e.Favorite_Sport = favoriteSport
}
