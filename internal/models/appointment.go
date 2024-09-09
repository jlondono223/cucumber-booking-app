package models

import "time"

type Appointment struct {
	AppointmentID	int 		`json:"appointment_id"`
	ClientID 		int 		`json:"client_id"`
	ProviderID 		int 		`json:"provider_id"`
	LocationID 		int 		`json:"location_id"`
	AppointmentDate time.Time 	`json:"appointment_date"`
	StartTime 		time.Time 	`json:"start_time"`
	EndTime 		time.Time 	`json:"end_time"`
	Status 			string 		`json:"status"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}