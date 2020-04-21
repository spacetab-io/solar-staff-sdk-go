package models

type CountryCode string

type PhoneNumber string

type WorkerCategory struct {
	ID uint64
	Title
}

type Worker struct {
	ID         uint64
	Email      string
	FirstName  string
	LastName   string
	Country    CountryCode
	Phone      *PhoneNumber
	Category   WorkerCategory `json:"workersCategory"`
	HasAddress bool
	Language   *CountryCode
	*Specialization
}
