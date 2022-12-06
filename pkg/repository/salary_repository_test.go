package repository_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"salaries/pkg/db"
	"salaries/pkg/domain"
	"salaries/pkg/repository"
	"testing"
)

type repositoryFields struct {
	dbClient db.DataBaseSalaryClient
}

func TestRepository_Create(t *testing.T) {
	tests := []struct {
		name      string
		salary    *domain.Salary
		fields    repositoryFields
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
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			salaryRepository := repository.NewSalaryRepositoryWithClient(tt.fields.dbClient)

			salary, err := salaryRepository.Create(tt.salary)
			if !tt.wantError {
				tt.salary.ID = salary.ID
				assert.Equal(t, tt.salary, salary)
			}
			assert.Equal(t, tt.wantError, err != nil)

		})
	}
}

func TestRepository_ReadAll(t *testing.T) {
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
		fields    repositoryFields
		wantError bool
	}{
		{
			name:     "success",
			salaries: salaries,
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			salaryRepository := repository.NewSalaryRepositoryWithClient(tt.fields.dbClient)

			salaries, err := salaryRepository.ReadAll()
			if !tt.wantError {
				assert.EqualValues(t, tt.salaries, salaries)
			}
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestRepository_DeleteByID(t *testing.T) {
	tests := []struct {
		name      string
		ID        int64
		fields    repositoryFields
		wantError bool
	}{
		{
			name: "success",
			ID:   1,
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			fields: repositoryFields{
				dbClient: &db.DataBaseSalaryClientMock{
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
			salaryRepository := repository.NewSalaryRepositoryWithClient(tt.fields.dbClient)

			err := salaryRepository.DeleteByID(tt.ID)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}
