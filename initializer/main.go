package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"os"
	dbClient "salaries/pkg/db"
	"salaries/pkg/domain"
	"salaries/pkg/logger"
)

const datasetPath = "../salaries/initializer/dataset.json"

// TODO: this is only to show how to add the dataset in the database to make easier and faster testing stats with endpoints
func main() {
	logger := logger.NewLogger()
	db, err := sql.Open("sqlite3", "salaries.db")
	if err != nil {
		logger.Error("error opening database")
	}

	createTable(db, logger)
	initializeSalaries(db, logger)
}

func createTable(dbClient *sql.DB, logger logger.Logger) error {
	statement, err := dbClient.Prepare("CREATE TABLE IF NOT EXISTS salaries (id INTEGER PRIMARY KEY, name VARCHAR(256), salary REAL, currency VARCHAR(64), on_contract INTEGER NULL, department VARCHAR(256), sub_department VARCHAR(256))")
	if err != nil {
		logger.Error("error creating table")
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		logger.Error("error creating table")
		return err
	}
	logger.Info("table salaries created")

	return nil
}

func initializeSalaries(dbClient *sql.DB, logger logger.Logger) {
	if !hasData(dbClient) {
		addSalariesFromJson(dbClient, logger)
	}
}

func hasData(dbClient *sql.DB) bool {
	statement, _ := dbClient.Prepare("SELECT COUNT(*) as count FROM salaries")
	var count int
	statement.QueryRow().Scan(&count)
	return count > 0
}

func addSalariesFromJson(dbClientSqlite *sql.DB, logger logger.Logger) {
	var data []byte
	data, err := os.ReadFile(datasetPath)
	if err != nil {
		logger.Info("Error reading file")
	}
	var salaries []domain.Salary
	json.Unmarshal(data, &salaries)
	dbClient := dbClient.NewSqlite(dbClientSqlite)
	for _, salary := range salaries {
		dbClient.Create(&salary)
	}
	logger.Info("Adding %d salaries from dataset", len(salaries))
}
