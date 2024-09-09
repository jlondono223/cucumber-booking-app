package handlers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "github.com/jlondono223/cucumber-booking-app/internal/repositories"
)

type ClientHandler struct {
    Repo *repositories.ClientRepository
}

func NewClientHandler(repo *repositories.ClientRepository) *ClientHandler {
    return &ClientHandler{Repo: repo}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := h.Repo.CreateClient(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client"})
        return
    }

    c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) GetClient(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
        return
    }

    client, err := h.Repo.GetClient(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
        return
    }

    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
    client.ClientID = id

    if err := h.Repo.UpdateClient(&client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client"})
        return
    }

    c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
        return
    }

    if err := h.Repo.DeleteClient(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h *ClientHandler) GetAllClients(c *gin.Context) {
    clients, err := h.Repo.GetAllClients()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clients"})
        return
    }

    c.JSON(http.StatusOK, clients)
}
