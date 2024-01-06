package httpExceptions

import (
	"net/http"
)

type HttpException struct {
	StatusCode    int    `json:"status"`
	Error         string `json:"message"`
	FriendlyError string `json:"friendlyMessage"`
}

func (err HttpException) IsEmpty() bool {
	return err.StatusCode == 0
}

var (
	EmployeeNotFound      = HttpException{http.StatusNotFound, "Employee not found.", "Could not find the desired employee with provided ID."}
	EmployeeAlreadyExists = HttpException{http.StatusNotFound, "Duplicate employee.", "Employee with provided ID already exists."}
	EmployeeIncomplete    = HttpException{http.StatusBadRequest, "Employee has missing fields.", "The submitted employee has missing fields."}
	InternalError         = HttpException{http.StatusInternalServerError, "Internal server error occured.", "Could not process your request at the moment."}
	MissingIDParameter    = HttpException{http.StatusBadRequest, "Missing ID parameter.", "Missing ID parameter."}
)
