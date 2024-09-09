package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // postgresql driver
)

var DB *sql.DB

type Config struct {
	Host		string
	Port		int
	User		string
	Password	string
	DBName		string
	SSLMode		string
}

func NewConfig() *Config {
	return &Config {
		Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnvAsInt("DB_PORT", 5432),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "bellaisboss23"),
        DBName:   getEnv("DB_NAME", "BABOOP_BACKEND"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func Connect(cfg *Config) error {
    // Data Source Name (DSN)
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("error opening database: %v", err)
    }

    // Ping the database to ensure the connection is established
    if err := DB.Ping(); err != nil {
        return fmt.Errorf("error connecting to the database: %v", err)
    }

    log.Println("Database connection established")
    return nil
}

// Close closes the database connection
func Close() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}

// Utility functions for reading environment variables

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
    if valueStr, exists := os.LookupEnv(name); exists {
        var value int
        _, err := fmt.Sscanf(valueStr, "%d", &value)
        if err == nil {
            return value
        }
    }
    return defaultValue
}