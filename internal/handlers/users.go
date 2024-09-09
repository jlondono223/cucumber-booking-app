package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "github.com/jlondono223/cucumber-booking-app/internal/repositories"
)


type UserHandler struct {
    Repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
    return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := h.Repo.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    users, err := h.Repo.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    c.JSON(http.StatusOK, users)
}