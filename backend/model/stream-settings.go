package model

type StreamSettings struct {	
	ID 				string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Title 			string `json:"title"`
	UserID 			string `json:"-"`
}
