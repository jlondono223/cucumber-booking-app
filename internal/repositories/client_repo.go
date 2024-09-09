package repositories

import (
    "database/sql"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "log"
)

type ClientRepository struct {
    DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
    return &ClientRepository{DB: db}
}

func (r *ClientRepository) CreateClient(client *models.Client) error {
    query := `
    INSERT INTO clients (user_id, address, dob, created_at, updated_at)
    VALUES ($1, $2, $3, NOW(), NOW()) RETURNING client_id`

    err := r.DB.QueryRow(query, client.UserID, client.Address, client.DOB).Scan(&client.ClientID)
    if err != nil {
        log.Println("Error creating client:", err)
        return err
    }
    return nil
}

func (r *ClientRepository) GetClient(id int) (*models.Client, error) {
    var client models.Client
    query := `SELECT client_id, user_id, address, dob, created_at, updated_at FROM clients WHERE client_id = $1`
    
    err := r.DB.QueryRow(query, id).Scan(
        &client.ClientID, &client.UserID, &client.Address, &client.DOB,
        &client.CreatedAt, &client.UpdatedAt)
    
    if err != nil {
        log.Println("Error getting client:", err)
        return nil, err
    }
    return &client, nil
}

func (r *ClientRepository) UpdateClient(client *models.Client) error {
    query := `
    UPDATE clients SET user_id=$1, address=$2, dob=$3, updated_at=NOW() WHERE client_id=$4`

    _, err := r.DB.Exec(query, client.UserID, client.Address, client.DOB, client.ClientID)
    if err != nil {
        log.Println("Error updating client:", err)
        return err
    }
    return nil
}

func (r *ClientRepository) DeleteClient(id int) error {
    query := `DELETE FROM clients WHERE client_id = $1`
    
    _, err := r.DB.Exec(query, id)
    if err != nil {
        log.Println("Error deleting client:", err)
        return err
    }
    return nil
}

func (r *ClientRepository) GetAllClients() ([]models.Client, error) {
    rows, err := r.DB.Query(`SELECT client_id, user_id, address, dob, created_at, updated_at FROM clients`)
    if err != nil {
        log.Println("Error getting clients:", err)
        return nil, err
    }
    defer rows.Close()

    var clients []models.Client
    for rows.Next() {
        var client models.Client
        err := rows.Scan(
            &client.ClientID, &client.UserID, &client.Address, &client.DOB,
            &client.CreatedAt, &client.UpdatedAt)
        
        if err != nil {
            log.Println("Error scanning client:", err)
            return nil, err
        }
        clients = append(clients, client)
    }
    return clients, nil
}
