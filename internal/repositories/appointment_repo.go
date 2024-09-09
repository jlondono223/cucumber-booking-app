package repositories

import (
    "database/sql"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "log"
)

type AppointmentRepository struct {
    DB *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
    return &AppointmentRepository{DB: db}
}

// CreateAppointment inserts a new appointment into the database
func (r *AppointmentRepository) CreateAppointment(appointment *models.Appointment) error {
    query := `
    INSERT INTO appointments (client_id, provider_id, location_id, appointment_date, start_time, end_time, status, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING appointment_id`

    err := r.DB.QueryRow(query, appointment.ClientID, appointment.ProviderID, appointment.LocationID, appointment.AppointmentDate, appointment.StartTime, appointment.EndTime, appointment.Status).Scan(&appointment.AppointmentID)
    if err != nil {
        log.Println("Error creating appointment:", err)
        return err
    }
    return nil
}

// GetAppointment retrieves an appointment by ID
func (r *AppointmentRepository) GetAppointment(id int) (*models.Appointment, error) {
    var appointment models.Appointment
    query := `SELECT appointment_id, client_id, provider_id, location_id, appointment_date, start_time, end_time, status, created_at, updated_at FROM appointments WHERE appointment_id = $1`
    
    err := r.DB.QueryRow(query, id).Scan(
        &appointment.AppointmentID, &appointment.ClientID, &appointment.ProviderID, &appointment.LocationID, 
        &appointment.AppointmentDate, &appointment.StartTime, &appointment.EndTime, &appointment.Status, 
        &appointment.CreatedAt, &appointment.UpdatedAt)
    
    if err != nil {
        log.Println("Error getting appointment:", err)
        return nil, err
    }
    return &appointment, nil
}

// UpdateAppointment updates an existing appointment in the database
func (r *AppointmentRepository) UpdateAppointment(appointment *models.Appointment) error {
    query := `
    UPDATE appointments SET client_id=$1, provider_id=$2, location_id=$3, appointment_date=$4, start_time=$5, end_time=$6, status=$7, updated_at=NOW() WHERE appointment_id=$8`

    _, err := r.DB.Exec(query, appointment.ClientID, appointment.ProviderID, appointment.LocationID, appointment.AppointmentDate, appointment.StartTime, appointment.EndTime, appointment.Status, appointment.AppointmentID)
    if err != nil {
        log.Println("Error updating appointment:", err)
        return err
    }
    return nil
}

// DeleteAppointment deletes an appointment by ID
func (r *AppointmentRepository) DeleteAppointment(id int) error {
    query := `DELETE FROM appointments WHERE appointment_id = $1`
    
    _, err := r.DB.Exec(query, id)
    if err != nil {
        log.Println("Error deleting appointment:", err)
        return err
    }
    return nil
}

// GetAllAppointments retrieves all appointments from the database
func (r *AppointmentRepository) GetAllAppointments() ([]models.Appointment, error) {
    rows, err := r.DB.Query(`SELECT appointment_id, client_id, provider_id, location_id, appointment_date, start_time, end_time, status, created_at, updated_at FROM appointments`)
    if err != nil {
        log.Println("Error getting appointments:", err)
        return nil, err
    }
    defer rows.Close()

    var appointments []models.Appointment
    for rows.Next() {
        var appointment models.Appointment
        err := rows.Scan(
            &appointment.AppointmentID, &appointment.ClientID, &appointment.ProviderID, &appointment.LocationID, 
            &appointment.AppointmentDate, &appointment.StartTime, &appointment.EndTime, &appointment.Status, 
            &appointment.CreatedAt, &appointment.UpdatedAt)
        
        if err != nil {
            log.Println("Error scanning appointment:", err)
            return nil, err
        }
        appointments = append(appointments, appointment)
    }
    return appointments, nil
}
