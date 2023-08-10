package models

type Artist struct {
	Username   string
	Country    string `json:"country" `
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
