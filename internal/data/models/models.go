package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Name      string               `bson:"name"`
	Age       int                  `bson:"age"`
	CourseIDs []primitive.ObjectID `bson:"course_ids"`
}

type Course struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
