package models

type Efficiency struct {
	Efficiency_ID    int     `json:"efficiency_id"`
	Efficiency_Score float64 `json:"efficiency_score"`
	Favorite_Sport   string  `json:"favorite_excercise"`
	Athlete_ID       string  `json:"athlete_id"`
}