package repository

import (
	"database/sql"
	"fmt"
	"service-employee/config"
	"service-employee/model"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
}

type employeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db: db}
}

func (empRepo *employeeRepositoryImpl) Create(employee *model.Employee) error {
	query := "INSERT INTO employees (name) VALUES ($1)"
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	_, errExec := empRepo.db.ExecContext(ctx, query, &employee.Name)
	if errExec != nil {
		// return errors.New("error, creating employee")
		return fmt.Errorf("error, create employee %w", errExec)
	}
	return nil
}
