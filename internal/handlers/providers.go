package handlers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "github.com/jlondono223/cucumber-booking-app/internal/repositories"
)

type ProviderHandler struct {
    Repo *repositories.ProviderRepository
}

func NewProviderHandler(repo *repositories.ProviderRepository) *ProviderHandler {
    return &ProviderHandler{Repo: repo}
}

func (h *ProviderHandler) CreateProvider(c *gin.Context) {
    var provider models.Provider
    if err := c.ShouldBindJSON(&provider); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := h.Repo.CreateProvider(&provider); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create provider"})
        return
    }

    c.JSON(http.StatusCreated, provider)
}

func (h *ProviderHandler) GetProvider(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
        return
    }

    provider, err := h.Repo.GetProvider(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found"})
        return
    }

    c.JSON(http.StatusOK, provider)
}

func (h *ProviderHandler) UpdateProvider(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
        return
    }

    var provider models.Provider
    if err := c.ShouldBindJSON(&provider); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
    provider.ProviderID = id

    if err := h.Repo.UpdateProvider(&provider); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update provider"})
        return
    }

    c.JSON(http.StatusOK, provider)
}

func (h *ProviderHandler) DeleteProvider(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider ID"})
        return
    }

    if err := h.Repo.DeleteProvider(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete provider"})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h *ProviderHandler) GetAllProviders(c *gin.Context) {
    providers, err := h.Repo.GetAllProviders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve providers"})
        return
    }

    c.JSON(http.StatusOK, providers)
}
