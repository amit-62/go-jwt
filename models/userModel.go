package models

import (
	"time"
)

type User struct {
	ID            primitive.ObjectID
	First_name    *string
	Last_name     *string
	Password      *string
	Email         *string
	Phone         *string
	Token         *string
	User_type     *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       string
}
