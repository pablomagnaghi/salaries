package main

import (
	"github.com/gin-gonic/gin"
	"salaries/pkg/auth"
	"salaries/pkg/controller"
	"salaries/pkg/logger"
	"salaries/pkg/middleware"
	"salaries/pkg/repository"
	"salaries/pkg/service"
)

const (
	TrustedProxy = "192.168.1.2"
	Port         = ":8080"
)

func main() {
	logger := logger.NewLogger()
	salaryRepository, err := repository.NewSalaryRepository(logger)
	if err != nil {
		logger.Error("failed to create salary repository: ", err.Error())
	}

	authService := auth.NewAuthService(logger)

	salaryService := service.NewSalaryService(salaryRepository, logger)

	serveApplication(authService, salaryService, logger)
}

func serveApplication(authService auth.Service, salaryService service.SalaryService, logger logger.Logger) {
	router := gin.Default()
	router.SetTrustedProxies([]string{TrustedProxy})

	authController := auth.NewAuthController(authService)

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", authController.Login)

	salaryController := controller.NewSalaryController(salaryService)

	protectedRoutes := router.Group("/api/salaries")
	protectedRoutes.Use(middleware.NewAuthMiddleware(authService))
	protectedRoutes.GET("", salaryController.GetAll)
	protectedRoutes.POST("", salaryController.Create)
	protectedRoutes.DELETE("/:id", salaryController.Delete)
	protectedRoutes.GET("/stats", salaryController.GetStatisticsEntireDataset)
	protectedRoutes.GET("/stats/contracts", salaryController.GetContractsStats)
	protectedRoutes.GET("/stats/departments", salaryController.GetDepartmentsStats)
	protectedRoutes.GET("/stats/sub-departments", salaryController.GetSubDepartmentsStats)

	router.Run(Port)
	logger.Info("Server api-creator running on port %s", Port)
}
