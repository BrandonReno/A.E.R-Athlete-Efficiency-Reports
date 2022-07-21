package models

type Excercise struct {
	ID          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Sets        []*Set `json:"sets"`
}
