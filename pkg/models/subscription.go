package models

type Subscription struct {
	User   User
	Course Course
	Price  float64
}
