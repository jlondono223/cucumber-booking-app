package handlers

import (
    
    "net/http"
    "strconv"
	"github.com/gin-gonic/gin"
    "github.com/jlondono223/cucumber-booking-app/internal/models"
    "github.com/jlondono223/cucumber-booking-app/internal/repositories"
    
)

type AppointmentHandler struct {
    Repo *repositories.AppointmentRepository
}

func NewAppointmentHandler(repo *repositories.AppointmentRepository) *AppointmentHandler {
    return &AppointmentHandler{Repo: repo}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := h.Repo.CreateAppointment(&appointment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
        return
    }

    c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
        return
    }

    appointment, err := h.Repo.GetAppointment(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
        return
    }

    c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
        return
    }

    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
    appointment.AppointmentID = id

    if err := h.Repo.UpdateAppointment(&appointment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
        return
    }

    c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
        return
    }

    if err := h.Repo.DeleteAppointment(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
    appointments, err := h.Repo.GetAllAppointments()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve appointments"})
        return
    }

    c.JSON(http.StatusOK, appointments)
}