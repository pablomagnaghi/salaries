package service_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"salaries/pkg/domain"
	"salaries/pkg/logger"
	"salaries/pkg/repository"
	"salaries/pkg/service"
	"testing"
)

type serviceFields struct {
	repository repository.SalaryRepositoryMock
}

func getTestLogger() logger.Logger {
	testLogger := logger.NewLogger()
	return testLogger
}

func TestRepository_Create(t *testing.T) {
	tests := []struct {
		name      string
		salary    *domain.Salary
		fields    serviceFields
		wantError bool
	}{
		{
			name: "success",
			salary: &domain.Salary{
				Name:          "Anurag",
				Salary:        90000,
				Currency:      "USD",
				Department:    "Banking",
				SubDepartment: "Loan",
			},
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					CreateFunc: func(salary *domain.Salary) (*domain.Salary, error) {
						return &domain.Salary{
							ID:            1,
							Name:          "Anurag",
							Salary:        90000,
							Currency:      "USD",
							Department:    "Banking",
							SubDepartment: "Loan",
						}, nil
					},
				},
			},
			wantError: false,
		},
		{
			name: "error",
			salary: &domain.Salary{
				Name:          "Anurag",
				Salary:        90000,
				Currency:      "USD",
				Department:    "Banking",
				SubDepartment: "Loan",
			},
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					CreateFunc: func(salary *domain.Salary) (*domain.Salary, error) {
						return nil, errors.New("error")
					},
				},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryService := service.NewSalaryService(&tt.fields.repository, getTestLogger())
			err := salaryService.Create(tt.salary)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
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
	tests := []struct {
		name      string
		salaries  []domain.Salary
		fields    serviceFields
		wantError bool
	}{
		{
			name:     "success",
			salaries: salaries,
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					ReadAllFunc: func() ([]domain.Salary, error) {
						return salaries, nil
					},
				},
			},
			wantError: false,
		},
		{
			name:     "error",
			salaries: nil,
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					ReadAllFunc: func() ([]domain.Salary, error) {
						return nil, errors.New("error")
					},
				},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryService := service.NewSalaryService(&tt.fields.repository, getTestLogger())

			salaries, err := salaryService.GetAll()
			if !tt.wantError {
				assert.EqualValues(t, tt.salaries, salaries)
			}
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	tests := []struct {
		name      string
		ID        int64
		fields    serviceFields
		wantError bool
	}{
		{
			name: "success",
			ID:   1,
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					DeleteByIDFunc: func(id int64) error {
						return nil
					},
				},
			},
			wantError: false,
		},
		{
			name: "error",
			ID:   2,
			fields: serviceFields{
				repository: repository.SalaryRepositoryMock{
					DeleteByIDFunc: func(id int64) error {
						return errors.New("error")
					},
				},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salaryService := service.NewSalaryService(&tt.fields.repository, getTestLogger())

			err := salaryService.DeleteByID(tt.ID)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}
