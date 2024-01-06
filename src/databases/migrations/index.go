package main

import (
	"github.com/greffgreff/gin-fiddle/src/databases"
	"github.com/greffgreff/gin-fiddle/src/models"
)

func init() {
	databases.ConnectToDB()
}

func main() {
	databases.DB.AutoMigrate(&models.Employee{})
}
