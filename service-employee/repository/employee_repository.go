package repository

import (
	"database/sql"
	"fmt"
	"service-employee/config"
	"service-employee/model"
)

type EmployeeRepository interface {
	Create(employee *model.Employee, db *sql.DB) error
}

type employeeRepositoryImpl struct {
}

func NewEmployeeRepositoryImpl() EmployeeRepository {
	return &employeeRepositoryImpl{}
}

func (empRepo *employeeRepositoryImpl) Create(employee *model.Employee, db *sql.DB) error {
	query := "INSERT INTO employees (name) VALUES ($1)"
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	_, errExec := db.ExecContext(ctx, query, &employee.Name)
	if errExec != nil {
		// return errors.New("error, creating employee")
		return fmt.Errorf("error, create employee %w", errExec)
	}
	return nil
}
