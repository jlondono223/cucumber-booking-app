package repositories

import (
    "database/sql"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "log"
)

type ProviderRepository struct {
    DB *sql.DB
}

func NewProviderRepository(db *sql.DB) *ProviderRepository {
    return &ProviderRepository{DB: db}
}

func (r *ProviderRepository) CreateProvider(provider *models.Provider) error {
    query := `
    INSERT INTO providers (user_id, profession, bio, created_at, updated_at)
    VALUES ($1, $2, $3, NOW(), NOW()) RETURNING provider_id`

    err := r.DB.QueryRow(query, provider.UserID, provider.Profession, provider.Bio).Scan(&provider.ProviderID)
    if err != nil {
        log.Println("Error creating provider:", err)
        return err
    }
    return nil
}

func (r *ProviderRepository) GetProvider(id int) (*models.Provider, error) {
    var provider models.Provider
    query := `SELECT provider_id, user_id, profession, bio, created_at, updated_at FROM providers WHERE provider_id = $1`
    
    err := r.DB.QueryRow(query, id).Scan(
        &provider.ProviderID, &provider.UserID, &provider.Profession, &provider.Bio,
        &provider.CreatedAt, &provider.UpdatedAt)
    
    if err != nil {
        log.Println("Error getting provider:", err)
        return nil, err
    }
    return &provider, nil
}

func (r *ProviderRepository) UpdateProvider(provider *models.Provider) error {
    query := `
    UPDATE providers SET user_id=$1, profession=$2, bio=$3, updated_at=NOW() WHERE provider_id=$4`

    _, err := r.DB.Exec(query, provider.UserID, provider.Profession, provider.Bio, provider.ProviderID)
    if err != nil {
        log.Println("Error updating provider:", err)
        return err
    }
    return nil
}

func (r *ProviderRepository) DeleteProvider(id int) error {
    query := `DELETE FROM providers WHERE provider_id = $1`
    
    _, err := r.DB.Exec(query, id)
    if err != nil {
        log.Println("Error deleting provider:", err)
        return err
    }
    return nil
}

func (r *ProviderRepository) GetAllProviders() ([]models.Provider, error) {
    rows, err := r.DB.Query(`SELECT provider_id, user_id, profession, bio, created_at, updated_at FROM providers`)
    if err != nil {
        log.Println("Error getting providers:", err)
        return nil, err
    }
    defer rows.Close()

    var providers []models.Provider
    for rows.Next() {
        var provider models.Provider
        err := rows.Scan(
            &provider.ProviderID, &provider.UserID, &provider.Profession, &provider.Bio,
            &provider.CreatedAt, &provider.UpdatedAt)
        
        if err != nil {
            log.Println("Error scanning provider:", err)
            return nil, err
        }
        providers = append(providers, provider)
    }
    return providers, nil
}
