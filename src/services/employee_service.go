package services

import (
	"errors"

	"github.com/greffgreff/gin-fiddle/src/databases"
	"github.com/greffgreff/gin-fiddle/src/httpExceptions"
	"github.com/greffgreff/gin-fiddle/src/models"
	"gorm.io/gorm"
)

func GetEmployeeList() []models.Employee {
	var employees []models.Employee
	databases.DB.Find(&employees)
	return employees
}

func GetEmployeeById(id string) (models.Employee, httpExceptions.HttpException) {
	var err httpExceptions.HttpException
	var employee models.Employee

	results := databases.DB.First(&employee, id)

	if errors.Is(results.Error, gorm.ErrRecordNotFound) {
		err = httpExceptions.EmployeeNotFound
	}

	return employee, err
}

func CreateEmployee(employee *models.Employee) *httpExceptions.HttpException {
	if employee.FirstName == "" || employee.LastName == "" {
		return &httpExceptions.EmployeeIncomplete
	}

	results := databases.DB.Create(&employee)

	if errors.Is(results.Error, gorm.ErrDuplicatedKey) {
		return &httpExceptions.EmployeeAlreadyExists
	}

	return nil
}

func UpdateEmployee(id string, employee *models.Employee) *httpExceptions.HttpException {
	// make use of gin validation mechanism instead
	if id == "" {
		return &httpExceptions.MissingIDParameter
	} else if employee.FirstName == "" || employee.LastName == "" {
		return &httpExceptions.EmployeeIncomplete
	}

	results := databases.DB.Updates(employee)

	if errors.Is(results.Error, gorm.ErrRecordNotFound) {
		return &httpExceptions.EmployeeNotFound
	}

	return nil
}

func DeleteEmployee(id string) *httpExceptions.HttpException {
	if id == "" {
		return &httpExceptions.MissingIDParameter
	}

	results := databases.DB.Delete(&models.Employee{}, id)

	if errors.Is(results.Error, gorm.ErrRecordNotFound) {
		return &httpExceptions.EmployeeNotFound
	}

	return nil
}
