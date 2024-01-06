package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greffgreff/gin-fiddle/src/httpExceptions"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Fatal("Unexpected error occured: " + err.Error())
		// Send error to error monitoring service
	}

	c.JSON(http.StatusInternalServerError, httpExceptions.InternalError)
}
