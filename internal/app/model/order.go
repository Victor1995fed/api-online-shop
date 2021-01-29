package model

//Order ...
type Order struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageURL    string `json:""`
	DateCreate  string `json:""`
}
