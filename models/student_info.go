package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Matric_No    int                `json:"matric_no,omitempty"`
	FirstName    string             `json:"first_name,omitempty"`
	LastName     string             `json:"last_name,omitempty"`
	Faculty      string             `json:"faculty,omitempty"`
	Department   string             `json:"department,omitempty"`
	Age          int                `json:"age,omitempty"`
	CGPA         float64            `json:"cgpa,omitempty"`
	InsertedTime time.Time          `json:"time"`
}
