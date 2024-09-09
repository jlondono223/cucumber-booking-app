package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jlondono223/cucumber-booking-app/internal/database"
    "github.com/jlondono223/cucumber-booking-app/internal/handlers"
    "github.com/jlondono223/cucumber-booking-app/internal/repositories"
)

func main() {
    // Initialize database connection
    dbConfig := database.NewConfig()
    if err := database.Connect(dbConfig); err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer database.Close()

    // Initialize repositories
    userRepo := repositories.NewUserRepository(database.DB)
    appointmentRepo := repositories.NewAppointmentRepository(database.DB)
    clientRepo := repositories.NewClientRepository(database.DB)
    providerRepo := repositories.NewProviderRepository(database.DB)

    // Initialize handlers
    userHandler := handlers.NewUserHandler(userRepo)
    appointmentHandler := handlers.NewAppointmentHandler(appointmentRepo)
    clientHandler := handlers.NewClientHandler(clientRepo)
    providerHandler := handlers.NewProviderHandler(providerRepo)

    // Create Gin router
    router := gin.Default()

    // Register user routes
    router.GET("/users", userHandler.GetUsers)
    router.POST("/users", userHandler.CreateUser)

    // Register appointment routes
    router.GET("/appointments", appointmentHandler.GetAllAppointments)
    router.POST("/appointments", appointmentHandler.CreateAppointment)
    router.GET("/appointments/:id", appointmentHandler.GetAppointment)
    router.PUT("/appointments/:id", appointmentHandler.UpdateAppointment)
    router.DELETE("/appointments/:id", appointmentHandler.DeleteAppointment)

    // Register client routes
    router.GET("/clients", clientHandler.GetAllClients)
    router.POST("/clients", clientHandler.CreateClient)
    router.GET("/clients/:id", clientHandler.GetClient)
    router.PUT("/clients/:id", clientHandler.UpdateClient)
    router.DELETE("/clients/:id", clientHandler.DeleteClient)

    // Register provider routes
    router.GET("/providers", providerHandler.GetAllProviders)
    router.POST("/providers", providerHandler.CreateProvider)
    router.GET("/providers/:id", providerHandler.GetProvider)
    router.PUT("/providers/:id", providerHandler.UpdateProvider)
    router.DELETE("/providers/:id", providerHandler.DeleteProvider)

    // Start the server
    go func() {
        if err := router.Run(":8080"); err != nil {
            log.Fatalf("Could not start server: %v\n", err)
        }
    }()

    // Handle graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    // Wait a few seconds to ensure server has properly shut down
    time.Sleep(2 * time.Second)
}
