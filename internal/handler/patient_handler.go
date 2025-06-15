package handler

import (
	"ass4/internal/models"
	"ass4/internal/utils"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	db *gorm.DB
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
	return &PatientHandler{db}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	claims, _ := utils.GetUserFromContext(c)
	patient.UpdatedBy = claims.UserID
	if err := h.db.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var patient models.Patient
	if err := h.db.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var req models.Patient
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	claims, _ := utils.GetUserFromContext(c)
	if claims.Role != "doctor" && claims.Role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}
	var patient models.Patient
	if err := h.db.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	patient.FirstName = req.FirstName
	patient.LastName = req.LastName
	patient.Age = req.Age
	patient.Gender = req.Gender
	patient.Address = req.Address
	patient.Phone = req.Phone
	patient.Details = req.Details
	patient.UpdatedBy = claims.UserID

	if err := h.db.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.db.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}

func (h *PatientHandler) ListPatients(c *gin.Context) {
	var patients []models.Patient
	if err := h.db.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}
