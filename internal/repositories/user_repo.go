package repositories

import (
	"database/sql"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}


func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, first_name, last_name, phone_number, role, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING user_id`

	err := r.DB.QueryRow(query, user.Email, user.Password, user.FirstName, user.LastName, user.PhoneNumber, user.Role).Scan(&user.UserID)
	if err != nil {
		log.Println("Error created user:", err)
		return err
	}
	return nil

}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    rows, err := r.DB.Query("SELECT user_id, email, first_name, last_name, phone_number, role, created_at, updated_at FROM users")
    if err != nil {
        log.Println("Error fetching users:", err)
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.UserID, &user.Email, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
            log.Println("Error scanning user:", err)
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}