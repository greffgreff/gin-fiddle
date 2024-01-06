package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greffgreff/gin-fiddle/src/models"
	"github.com/greffgreff/gin-fiddle/src/services"
)

func InitializeEmployeeController(rg *gin.RouterGroup) {
	employeeRouter := rg.Group("/employees")
	{
		employeeRouter.GET("", getEmployeeList)
		employeeRouter.GET("/:id", getEmployeeById)
		employeeRouter.POST("", createEmployee)
		employeeRouter.PUT("/:id", updateEmployee)
		employeeRouter.DELETE("/:id", deleteEmployee)
	}
}

func getEmployeeList(c *gin.Context) {
	employees := services.GetEmployeeList()
	c.JSON(http.StatusOK, employees)
}

func getEmployeeById(c *gin.Context) {
	id := c.Param("id")

	employee, err := services.GetEmployeeById(id)

	if !err.IsEmpty() {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func createEmployee(c *gin.Context) {
	var employee models.Employee

	c.Bind(&employee)

	err := services.CreateEmployee(&employee)

	if !err.IsEmpty() {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	c.Bind(&employee)

	err := services.UpdateEmployee(id, &employee)

	if !err.IsEmpty() {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteEmployee(id)

	if !err.IsEmpty() {
		c.JSON(err.StatusCode, err)
		return
	}

	c.Status(http.StatusOK)
}
