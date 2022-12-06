package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"salaries/pkg/api"
	dbClient "salaries/pkg/db"
	"salaries/pkg/domain"
	"salaries/pkg/logger"
)

type SalaryRepository interface {
	Create(salary *domain.Salary) (*domain.Salary, error)
	ReadAll() ([]domain.Salary, error)
	DeleteByID(salaryID int64) error
	GetStatsForAllSalaries() (*api.Stats, error)
	GetContractsStats() (*api.Stats, error)
	GetDepartmentsStats() ([]api.DepartmentStats, error)
	GetSubDepartmentsStats() ([]api.SubDepartmentStats, error)
}

type salaryRepositoryImpl struct {
	dbClient dbClient.DataBaseSalaryClient
}

func NewSalaryRepository(logger logger.Logger) (SalaryRepository, error) {
	db, err := sql.Open("sqlite3", "salaries.db")

	if err != nil {
		logger.Error("error open database", err)
	}

	dbClient := dbClient.NewSqlite(db)
	return NewSalaryRepositoryWithClient(dbClient), nil
}

func NewSalaryRepositoryWithClient(dbClient dbClient.DataBaseSalaryClient) SalaryRepository {
	return &salaryRepositoryImpl{
		dbClient: dbClient,
	}
}

func (s salaryRepositoryImpl) Create(salary *domain.Salary) (*domain.Salary, error) {
	return s.dbClient.Create(salary)
}

func (s salaryRepositoryImpl) ReadAll() ([]domain.Salary, error) {
	return s.dbClient.ReadAll()
}

func (s salaryRepositoryImpl) DeleteByID(salaryID int64) error {
	return s.dbClient.DeleteByID(salaryID)
}

func (s salaryRepositoryImpl) GetStatsForAllSalaries() (*api.Stats, error) {
	return s.dbClient.GetStatsForAllSalaries()
}

func (s salaryRepositoryImpl) GetContractsStats() (*api.Stats, error) {
	return s.dbClient.GetContractsStats()
}

func (s salaryRepositoryImpl) GetDepartmentsStats() ([]api.DepartmentStats, error) {
	return s.dbClient.GetDepartmentsStats()
}

func (s salaryRepositoryImpl) GetSubDepartmentsStats() ([]api.SubDepartmentStats, error) {
	return s.dbClient.GetSubDepartmentsStats()
}
