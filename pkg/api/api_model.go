package api

type AppError struct {
	Error error
	Code  int
}

type Stats struct {
	Mean float64
	Max  float64
	Min  float64
}

type DepartmentStats struct {
	Department string
	Stats      Stats
}

type SubDepartmentStats struct {
	SubDepartment   string
	DepartmentStats DepartmentStats
}
