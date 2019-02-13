package models

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Cars      []Car  `json:"cars"`
}
