package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ass4/internal/handler"
	"ass4/internal/models"
	"ass4/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Patient{})
	return db
}

func setupRouterWithUser(db *gorm.DB, user *utils.Claims) *gin.Engine {
	h := handler.NewPatientHandler(db)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("user", user)
		c.Next()
	})
	r.POST("/patients", h.CreatePatient)
	r.GET("/patients", h.ListPatients)
	r.GET("/patients/:id", h.GetPatient)
	r.PUT("/patients/:id", h.UpdatePatient)
	r.DELETE("/patients/:id", h.DeletePatient)
	return r
}

func TestCreatePatientHandler(t *testing.T) {
	db := setupTestDB()
	r := setupRouterWithUser(db, &utils.Claims{UserID: 1, Role: "receptionist"})
	patient := map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
		"age":       30,
		"gender":    "male",
		"address":   "123 St",
		"phone":     "1234567890",
		"details":   "Healthy",
	}
	body, _ := json.Marshal(patient)
	req, _ := http.NewRequest("POST", "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListPatientsHandler(t *testing.T) {
	db := setupTestDB()
	// Seed a patient
	db.Create(&models.Patient{FirstName: "John", LastName: "Doe"})
	r := setupRouterWithUser(db, &utils.Claims{UserID: 1, Role: "doctor"})
	req, _ := http.NewRequest("GET", "/patients", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPatientHandler(t *testing.T) {
	db := setupTestDB()
	patient := models.Patient{FirstName: "John", LastName: "Doe"}
	db.Create(&patient)
	r := setupRouterWithUser(db, &utils.Claims{UserID: 1, Role: "doctor"})
	url := "/patients/1"
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePatientHandler(t *testing.T) {
	db := setupTestDB()
	patient := models.Patient{FirstName: "John", LastName: "Doe"}
	db.Create(&patient)
	r := setupRouterWithUser(db, &utils.Claims{UserID: 2, Role: "doctor"})
	update := map[string]interface{}{
		"firstName": "Jane",
		"lastName":  "Smith",
		"age":       40,
		"gender":    "female",
		"address":   "456 Ave",
		"phone":     "9876543210",
		"details":   "Updated",
	}
	body, _ := json.Marshal(update)
	req, _ := http.NewRequest("PUT", "/patients/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePatientHandler(t *testing.T) {
	db := setupTestDB()
	patient := models.Patient{FirstName: "John", LastName: "Doe"}
	db.Create(&patient)
	r := setupRouterWithUser(db, &utils.Claims{UserID: 1, Role: "receptionist"})
	req, _ := http.NewRequest("DELETE", "/patients/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
