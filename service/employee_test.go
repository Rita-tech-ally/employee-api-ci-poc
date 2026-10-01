package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository implementation using Testify Mock
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) GetByID(id string) (*Employee, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Save(emp *Employee) error {
	args := m.Called(emp)
	return args.Error(0)
}

// Test cases
func TestCreateEmployee_Success(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	emp := &Employee{ID: "EMP-01", Name: "Abhishek Dubey", Designation: "Consultant"}
	mockRepo.On("Save", emp).Return(nil)

	err := svc.CreateEmployee(emp)

	assert.NoError(t, err)
	assert.Equal(t, "Active Employee", emp.Status)
	mockRepo.AssertExpectations(t)
}

func TestCreateEmployee_EmptyID(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	emp := &Employee{ID: "", Name: "Abhishek Dubey"}
	err := svc.CreateEmployee(emp)

	assert.Error(t, err)
	assert.Equal(t, "employee id cannot be empty", err.Error())
}

func TestCreateEmployee_EmptyName(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	emp := &Employee{ID: "EMP-01", Name: ""}
	err := svc.CreateEmployee(emp)

	assert.Error(t, err)
	assert.Equal(t, "employee name cannot be empty", err.Error())
}

func TestGetEmployee_Success(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	expectedEmp := &Employee{ID: "EMP-01", Name: "Abhishek Dubey", Status: "Active Employee"}
	mockRepo.On("GetByID", "EMP-01").Return(expectedEmp, nil)

	result, err := svc.GetEmployee("EMP-01")

	assert.NoError(t, err)
	assert.Equal(t, "Abhishek Dubey", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetEmployee_NotFound(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	mockRepo.On("GetByID", "EMP-99").Return(nil, errors.New("employee not found"))

	result, err := svc.GetEmployee("EMP-99")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "employee not found", err.Error())
}

func TestGetEmployee_EmptyID(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	svc := NewEmployeeService(mockRepo)

	result, err := svc.GetEmployee("")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "invalid id provided", err.Error())
}
