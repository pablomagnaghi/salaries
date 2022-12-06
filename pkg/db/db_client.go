package db

import (
	"database/sql"
	"salaries/pkg/api"
	"salaries/pkg/domain"
)

type DataBaseSalaryClient interface {
	Create(salary *domain.Salary) (*domain.Salary, error)
	ReadAll() ([]domain.Salary, error)
	DeleteByID(salaryID int64) error
	GetStatsForAllSalaries() (*api.Stats, error)
	GetContractsStats() (*api.Stats, error)
	GetDepartmentsStats() ([]api.DepartmentStats, error)
	GetSubDepartmentsStats() ([]api.SubDepartmentStats, error)
}

func NewSqlite(client *sql.DB) DataBaseSalaryClient {
	return &dataBaseClientImpl{
		client: client,
	}
}

type dataBaseClientImpl struct {
	client *sql.DB
}

func (d dataBaseClientImpl) Create(salary *domain.Salary) (*domain.Salary, error) {
	statement, err := d.client.Prepare("INSERT INTO salaries (name, salary, currency, on_contract, department, sub_department) VALUES (?, ?, ?, ?, ?, ?)")
	result, err := statement.Exec(salary.Name, salary.Salary, salary.Currency, salary.OnContract, salary.Department, salary.SubDepartment)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	salary.ID = id
	return salary, nil
}

func (d dataBaseClientImpl) ReadAll() ([]domain.Salary, error) {
	rows, _ := d.client.Query("SELECT id, name, salary, currency, on_contract, department, sub_department FROM salaries")
	var salaries []domain.Salary
	var salary domain.Salary
	for rows.Next() {
		rows.Scan(&salary.ID, &salary.Name, &salary.Salary, &salary.Currency, &salary.OnContract, &salary.Department, &salary.SubDepartment)
		salaries = append(salaries, salary)
	}
	return salaries, nil
}

func (d dataBaseClientImpl) DeleteByID(salaryID int64) error {
	query := "delete from salaries where id=?"
	statement, err := d.client.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(salaryID)
	if err != nil {
		return err
	}
	return nil
}

func (d dataBaseClientImpl) GetStatsForAllSalaries() (*api.Stats, error) {
	rows, err := d.client.Query("SELECT MIN(salary) as min, MAX(salary) as max, AVG(salary) as mean FROM salaries")
	if err != nil {
		return nil, err
	}
	var min, max, mean float64
	for rows.Next() {
		rows.Scan(&min, &max, &mean)
	}
	return &api.Stats{
		Mean: mean,
		Max:  max,
		Min:  min,
	}, nil
}

func (d dataBaseClientImpl) GetContractsStats() (*api.Stats, error) {
	rows, err := d.client.Query("SELECT MIN(salary) as minPrice, MAX(salary) as maxPrice, AVG(salary) as mean FROM salaries WHERE on_contract = 1")
	if err != nil {
		return nil, err
	}
	var min, max, mean float64
	for rows.Next() {
		rows.Scan(&min, &max, &mean)
	}
	return &api.Stats{
		Mean: mean,
		Max:  max,
		Min:  min,
	}, nil
}

func (d dataBaseClientImpl) GetDepartmentsStats() ([]api.DepartmentStats, error) {
	rows, err := d.client.Query("SELECT department as subDepartment, MIN(salary) as minPrice, MAX(salary) as maxPrice, AVG(salary) as mean FROM salaries GROUP BY department")
	if err != nil {
		return nil, err
	}
	var departmentsStats []api.DepartmentStats
	var department string
	var min, max, mean float64
	for rows.Next() {
		rows.Scan(&department, &min, &max, &mean)
		departmentStats := api.DepartmentStats{
			Department: department,
			Stats: api.Stats{
				Mean: mean,
				Max:  max,
				Min:  min,
			},
		}
		departmentsStats = append(departmentsStats, departmentStats)
	}
	return departmentsStats, nil
}

func (d dataBaseClientImpl) GetSubDepartmentsStats() ([]api.SubDepartmentStats, error) {
	rows, err := d.client.Query("SELECT department, sub_department as subDepartment, MIN(salary) as minPrice, MAX(salary) as maxPrice, AVG(salary) as mean FROM salaries GROUP BY department, sub_department")
	if err != nil {
		return nil, err
	}
	var subDepartmentsStats []api.SubDepartmentStats
	var department, subDepartment string
	var min, max, mean float64
	for rows.Next() {
		rows.Scan(&department, &subDepartment, &min, &max, &mean)
		departmentStats := api.SubDepartmentStats{
			SubDepartment: subDepartment,
			DepartmentStats: api.DepartmentStats{
				Department: department,
				Stats: api.Stats{
					Mean: mean,
					Max:  max,
					Min:  min,
				},
			},
		}
		subDepartmentsStats = append(subDepartmentsStats, departmentStats)
	}
	return subDepartmentsStats, nil
}
