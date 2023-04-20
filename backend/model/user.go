package model

import (
	"time"
)

type User struct {	
	ID 				string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Username 		string `json:"username"`
	Password 		string `json:"password,omitempty"`
	IsAdmin 		bool `json:"is_admin"`
	StreamKey 		string `json:"stream_key,omitempty"`
	IsLive  		bool `json:"is_live"`
	LastStreamAt 	time.Time `json:"last_stream_at"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}
