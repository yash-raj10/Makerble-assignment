package routes

import (
	"ass4/internal/handler"
	"ass4/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handler.AuthHandler, patientHandler *handler.PatientHandler) {
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	patient := r.Group("/patients")
	patient.Use(middleware.AuthMiddleware())
	{
		patient.POST("", middleware.RoleMiddleware("receptionist"), patientHandler.CreatePatient)
		patient.GET("", patientHandler.ListPatients)
		patient.GET(":id", patientHandler.GetPatient)
		patient.PUT(":id", middleware.RoleMiddleware("doctor"), patientHandler.UpdatePatient)
		patient.DELETE(":id", middleware.RoleMiddleware("receptionist"), patientHandler.DeletePatient)
	}
}
