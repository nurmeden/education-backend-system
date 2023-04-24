package model

import "time"

type Student struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Password  string `bson:"password"`
	Age       int    `bson:"age"`
}

type SignInData struct {
	UserID   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthToken struct {
	UserID    string    `json:"userId"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}
