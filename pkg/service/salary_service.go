package service

import (
	"fmt"
	"salaries/pkg/api"
	"salaries/pkg/domain"
	"salaries/pkg/logger"
	"salaries/pkg/repository"
)

type SalaryService interface {
	Create(*domain.Salary) error
	GetAll() ([]domain.Salary, error)
	DeleteByID(id int64) error
	GetStatsForAllSalaries() (*api.Stats, error)
	GetContractsStats() (*api.Stats, error)
	GetDepartmentsStats() ([]api.DepartmentStats, error)
	GetSubDepartmentsStats() ([]api.SubDepartmentStats, error)
}

type salaryServiceImpl struct {
	salaryRepository repository.SalaryRepository
	logger           logger.Logger
}

func NewSalaryService(salaryRepository repository.SalaryRepository, logger logger.Logger) SalaryService {
	return &salaryServiceImpl{
		salaryRepository: salaryRepository,
		logger:           logger,
	}
}

func (s salaryServiceImpl) Create(salary *domain.Salary) error {
	s.logger.Info(fmt.Sprintf("creating salary %v", salary))
	salary, err := s.salaryRepository.Create(salary)
	if err != nil {
		return err
	}
	s.logger.Info(fmt.Sprintf("salary created %v", salary))
	return nil
}

func (s salaryServiceImpl) GetAll() ([]domain.Salary, error) {
	s.logger.Info("Getting all salaries")
	salaries, err := s.salaryRepository.ReadAll()
	if err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("salaries retrieved"))
	return salaries, nil
}

func (s salaryServiceImpl) DeleteByID(salaryID int64) error {
	s.logger.Info(fmt.Sprintf("deleting salary with id %d", salaryID))
	err := s.salaryRepository.DeleteByID(salaryID)
	if err != nil {
		return err
	}
	s.logger.Info(fmt.Sprintf("salary deleted with id %d", salaryID))
	return nil
}

func (s salaryServiceImpl) GetStatsForAllSalaries() (*api.Stats, error) {
	s.logger.Info("Getting stats")
	stats, err := s.salaryRepository.GetStatsForAllSalaries()
	if err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("stats retrieved"))
	return stats, nil
}

func (s salaryServiceImpl) GetContractsStats() (*api.Stats, error) {
	s.logger.Info("Getting contract stats")
	stats, err := s.salaryRepository.GetContractsStats()
	if err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("contract stats retrieved"))
	return stats, nil
}

func (s salaryServiceImpl) GetDepartmentsStats() ([]api.DepartmentStats, error) {
	s.logger.Info("Getting departments stats")
	stats, err := s.salaryRepository.GetDepartmentsStats()
	if err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("departments stats retrieved"))
	return stats, nil
}

func (s salaryServiceImpl) GetSubDepartmentsStats() ([]api.SubDepartmentStats, error) {
	s.logger.Info("Getting sub-departments stats")
	stats, err := s.salaryRepository.GetSubDepartmentsStats()
	if err != nil {
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("sub-departments stats retrieved"))
	return stats, nil
}
