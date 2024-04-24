package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type DueTime struct {
	Date   primitive.DateTime
	Repeat string
}

type Todo struct {
	Id    string
	Title string
	Notes string
	Due   DueTime
}
