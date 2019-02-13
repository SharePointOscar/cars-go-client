package models

//Car Model
type Car struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
	Year  int64  `json:"year"`
	Make  string `json:"make"`
}

//Person Model
type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Cars      []Car  `json:"cars"`
}
