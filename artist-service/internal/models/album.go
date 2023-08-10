package models

type Album struct {
	Title       string `json:"title"`
	Song        []Song `json:"songs"`
	ReleaseYear int    `json:"release_year"`
	Country     string `json:"country"`
	Cover       []byte `json:"cover"`
}
