package model

type StreamCredentials struct {	
	ID 				string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	RTMPUrl 		string `json:"rtmp_url"`
	StreamKey 		string `json:"stream_key"`
	UserID 			string `json:"user_id"`
}
