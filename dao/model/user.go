// contain User model
package model

import "time"

type User struct {
	// ID primary key
	ID    string `json:"id"`
	Email string `json:"email"`
	// default value is password
	Password    string    `json:"password" gorm:"default:'password'"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CreatedDate time.Time `json:"created_date"`
}
