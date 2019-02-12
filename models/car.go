package models

type Car struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Make  string `json:"genre"`
}
