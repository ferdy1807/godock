package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthdate"`
	Gender    string `json:"gender"`
	Job       string `json:"job"`
	Photo     string `json:"photo"`
}
