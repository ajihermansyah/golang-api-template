package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name      string        `bson:"name" json:"name"`
	Age       int           `bson:"age" json:"age"`
	Email     string        `bson:"email" json:"email"`
	Gender    string        `bson:"gender" json:"gender"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
