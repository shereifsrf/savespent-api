package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user sessoin with created date
type UserSession struct {
	// ID primary key
	ID primitive.ObjectID `json:"id" bson:"_id"`
	// user ID
	UserID string `json:"user_id" bson:"user_id"`
	// device type
	DeviceType string `json:"devide_type" bson:"devide_type"`
	// created date is a date type
	CreatedDate time.Time `json:"created_date" bson:"created_date"`
}
