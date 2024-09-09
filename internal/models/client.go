package models

import "time"

type Client struct {
	ClientID  int       `json:"client_id"`
    UserID    int       `json:"user_id"`
    Address   string    `json:"address"`
    DOB       time.Time `json:"dob"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}