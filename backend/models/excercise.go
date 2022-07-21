package models

type Excercise struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Sets        []*Set `json:"sets"`
}
