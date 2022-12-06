package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"salaries/pkg/api"
	"salaries/pkg/domain"
	"salaries/pkg/service"
	"strconv"
)

type SalaryController interface {
	Create(context *gin.Context)
	GetAll(context *gin.Context)
	Delete(context *gin.Context)
	GetStatisticsEntireDataset(context *gin.Context)
	GetContractsStats(context *gin.Context)
	GetDepartmentsStats(context *gin.Context)
	GetSubDepartmentsStats(context *gin.Context)
}

type salaryControllerImpl struct {
	salaryService service.SalaryService
}

func NewSalaryController(salaryService service.SalaryService) SalaryController {
	return &salaryControllerImpl{
		salaryService: salaryService,
	}
}

func (c salaryControllerImpl) Create(context *gin.Context) {
	var salary domain.Salary
	if err := context.ShouldBindJSON(&salary); err != nil {
		c.badRequest(context, err)
		return
	}

	err := c.salaryService.Create(&salary)
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusCreated, salary)
}

func (c salaryControllerImpl) GetAll(context *gin.Context) {
	salaries, err := c.salaryService.GetAll()
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusOK, salaries)
}

func (c salaryControllerImpl) Delete(context *gin.Context) {
	ID := context.Param("id")
	salaryID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		c.badRequest(context, err)
		return
	}

	appError := c.salaryService.DeleteByID(salaryID)
	if appError != nil {
		c.internalServerError(context, err)
		return
	}
	context.Status(http.StatusNoContent)
}

func (c salaryControllerImpl) GetStatisticsEntireDataset(context *gin.Context) {
	stats, err := c.salaryService.GetStatsForAllSalaries()
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusOK, stats)
}

func (c salaryControllerImpl) GetContractsStats(context *gin.Context) {
	stats, err := c.salaryService.GetContractsStats()
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusOK, stats)
}

func (c salaryControllerImpl) GetDepartmentsStats(context *gin.Context) {
	stats, err := c.salaryService.GetDepartmentsStats()
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusOK, stats)
}

func (c salaryControllerImpl) GetSubDepartmentsStats(context *gin.Context) {
	stats, err := c.salaryService.GetSubDepartmentsStats()
	if err != nil {
		c.internalServerError(context, err)
		return
	}
	context.JSON(http.StatusOK, stats)
}

func (c salaryControllerImpl) badRequest(context *gin.Context, err error) {
	context.JSON(http.StatusBadRequest, api.AppError{
		Error: err,
		Code:  http.StatusBadRequest,
	})
}

func (c salaryControllerImpl) internalServerError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, api.AppError{
		Error: err,
		Code:  http.StatusInternalServerError,
	})
}
