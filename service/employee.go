package service

import (
	"errors"
	"strings"
)

// Employee domain model
type Employee struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Designation    string `json:"designation"`
	Department     string `json:"department"`
	OfficeLocation string `json:"office_location"`
	Status         string `json:"status"`
}

// EmployeeRepository interface jise tests me mock kiya jayega
type EmployeeRepository interface {
	GetByID(id string) (*Employee, error)
	Save(emp *Employee) error
}

// EmployeeService business logic container
type EmployeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// CreateEmployee validates and saves employee
func (s *EmployeeService) CreateEmployee(emp *Employee) error {
	if strings.TrimSpace(emp.ID) == "" {
		return errors.New("employee id cannot be empty")
	}
	if strings.TrimSpace(emp.Name) == "" {
		return errors.New("employee name cannot be empty")
	}
	emp.Status = "Active Employee"
	return s.repo.Save(emp)
}

// GetEmployee fetches employee by ID
func (s *EmployeeService) GetEmployee(id string) (*Employee, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("invalid id provided")
	}
	return s.repo.GetByID(id)
}
