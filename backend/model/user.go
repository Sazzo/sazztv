package model

import (
	"time"
)

type User struct {	
	ID 				string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Username 		string `json:"username"`
	Password 		string `json:"password,omitempty"`
	IsAdmin 		bool `json:"is_admin"`
	StreamSettings 	StreamSettings `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"stream_settings,omitempty"`
	StreamCredentials StreamCredentials `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"stream_credentials,omitempty"`
	IsLive  		bool `json:"is_live"`
	LastStreamAt 	time.Time `json:"last_stream_at"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}