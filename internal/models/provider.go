package models

import "time"

type Provider struct {
    ProviderID int       `json:"provider_id"`
    UserID     int       `json:"user_id"`
    Profession string    `json:"profession"`
    Bio        string    `json:"bio"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
