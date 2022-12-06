package controller_test

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"salaries/pkg/api"
	"salaries/pkg/auth"
	"salaries/pkg/controller"
	"salaries/pkg/domain"
	"salaries/pkg/middleware"
	"salaries/pkg/service"
	"strings"
	"testing"
)

var authService = &auth.ServiceMock{
	VerifyTokenFunc: func(context *gin.Context) error {
		return nil
	},
}

func TestHTTPHandler_Create(t *testing.T) {
	salary := "{\"name\": \"Anurag\", \"salary\": \"90000\", \"currency\": \"USD\", " +
		"\"department\": \"Banking\", \"on_contract\": \"true\",  \"sub_department\": \"Loan\"}"
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		salary    string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name:   "invalid token",
			salary: salary,
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name:   "create salary",
			salary: salary,
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					CreateFunc: func(salary *domain.Salary) error {
						return nil
					},
				},
			},
			status:    http.StatusCreated,
			wantError: true,
		},
		{
			name:   "error creating salary",
			salary: salary,
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					CreateFunc: func(salary *domain.Salary) error {
						return errors.New("error creating salary")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.POST("/api/salaries", salaryController.Create)

			bodyReader := strings.NewReader(tt.salary)

			req := httptest.NewRequest(http.MethodPost, "/api/salaries", bodyReader)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_GetAll(t *testing.T) {
	salaries := []domain.Salary{
		{
			ID:            1,
			Name:          "Anurag",
			Salary:        90000,
			Currency:      "USD",
			Department:    "Banking",
			SubDepartment: "Loan",
		},
	}
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "get salaries",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetAllFunc: func() ([]domain.Salary, error) {
						return salaries, nil
					},
				},
			},
			status:    http.StatusOK,
			wantError: true,
		},
		{
			name: "error getting salaries",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetAllFunc: func() ([]domain.Salary, error) {
						return nil, errors.New("error getting salaries")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.GET("/api/salaries", salaryController.GetAll)

			req := httptest.NewRequest(http.MethodGet, "/api/salaries", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_Delete(t *testing.T) {
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		ID        int64
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			ID:   1,
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "delete salary",
			ID:   1,
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					DeleteByIDFunc: func(ID int64) error {
						return nil
					},
				},
			},
			status:    http.StatusNoContent,
			wantError: true,
		},
		{
			name: "error deleting salary",
			ID:   2,
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					DeleteByIDFunc: func(ID int64) error {
						return errors.New("error deleting salary")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.DELETE("/api/salaries/:id", salaryController.Delete)

			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/salaries/%d", tt.ID), nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_GetStatisticsEntireDataset(t *testing.T) {
	stats := api.Stats{
		Mean: 30.0,
		Min:  15.0,
		Max:  45.0,
	}
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "get stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetStatsForAllSalariesFunc: func() (*api.Stats, error) {
						return &stats, nil
					},
				},
			},
			status:    http.StatusOK,
			wantError: true,
		},
		{
			name: "error getting stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetStatsForAllSalariesFunc: func() (*api.Stats, error) {
						return nil, errors.New("error getting stats")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.GET("/api/salaries/stats", salaryController.GetStatisticsEntireDataset)

			req := httptest.NewRequest(http.MethodGet, "/api/salaries/stats", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_GetContractsStats(t *testing.T) {
	stats := api.Stats{
		Mean: 30.0,
		Min:  15.0,
		Max:  45.0,
	}
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "get contracts stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetContractsStatsFunc: func() (*api.Stats, error) {
						return &stats, nil
					},
				},
			},
			status:    http.StatusOK,
			wantError: true,
		},
		{
			name: "error getting contracts stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetContractsStatsFunc: func() (*api.Stats, error) {
						return nil, errors.New("error getting contracts stats")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.GET("/api/salaries/stats/contracts", salaryController.GetContractsStats)

			req := httptest.NewRequest(http.MethodGet, "/api/salaries/stats/contracts", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_GetDepartmentsStats(t *testing.T) {
	stats := []api.DepartmentStats{
		{
			Department: "Banking",
			Stats: api.Stats{
				Mean: 30.0,
				Min:  15.0,
				Max:  45.0,
			},
		},
	}
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "getting departments stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetDepartmentsStatsFunc: func() ([]api.DepartmentStats, error) {
						return stats, nil
					},
				},
			},
			status:    http.StatusOK,
			wantError: true,
		},
		{
			name: "error getting departments stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetDepartmentsStatsFunc: func() ([]api.DepartmentStats, error) {
						return nil, errors.New("error getting departments stats")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.GET("/api/salaries/stats/departments", salaryController.GetDepartmentsStats)

			req := httptest.NewRequest(http.MethodGet, "/api/salaries/stats/departments", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestHTTPHandler_GetSubDepartmentsStats(t *testing.T) {
	stats := []api.SubDepartmentStats{
		{
			SubDepartment: "Loan",
			DepartmentStats: api.DepartmentStats{
				Department: "Banking",
				Stats: api.Stats{
					Mean: 30.0,
					Min:  15.0,
					Max:  45.0,
				},
			}},
	}
	type fields struct {
		authService   auth.Service
		salaryService service.SalaryService
	}
	tests := []struct {
		name      string
		fields    fields
		status    int
		wantError bool
	}{
		{
			name: "invalid token",
			fields: fields{
				authService: &auth.ServiceMock{
					VerifyTokenFunc: func(context *gin.Context) error {
						return errors.New("invalid token")
					},
				},
			},
			status:    http.StatusUnauthorized,
			wantError: true,
		},
		{
			name: "getting subDepartments stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetSubDepartmentsStatsFunc: func() ([]api.SubDepartmentStats, error) {
						return stats, nil
					},
				},
			},
			status:    http.StatusOK,
			wantError: true,
		},
		{
			name: "error getting subDepartments stats",
			fields: fields{
				authService: authService,
				salaryService: &service.SalaryServiceMock{
					GetSubDepartmentsStatsFunc: func() ([]api.SubDepartmentStats, error) {
						return nil, errors.New("error getting subDepartments stats")
					},
				},
			},
			status:    http.StatusInternalServerError,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryController := controller.NewSalaryController(tt.fields.salaryService)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(middleware.NewAuthMiddleware(tt.fields.authService))

			r.GET("/api/salaries/stats/sub-departments", salaryController.GetSubDepartmentsStats)

			req := httptest.NewRequest(http.MethodGet, "/api/salaries/stats/sub-departments", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}
