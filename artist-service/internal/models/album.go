package models

type Album struct {
	Title string `json:"title"`
	Song  []Song `json:"songs"`
}
