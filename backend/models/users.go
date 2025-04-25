package models

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	BirthDate string `json:"birthdate"` // bisa juga time.Time kalau kamu ingin support parsing tanggal
	Gender    string `json:"gender"`
	Job       string `json:"job"`
	Photo     string `json:"photo"`
}
