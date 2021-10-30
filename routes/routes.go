package routes

import (
	"github.com/alexdang1993374/employee-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/employees", controllers.CreateEmployee)
	router.PUT("/employees/:employeeId", controllers.UpdateEmployee)
	router.DELETE("/employees/:employeeId", controllers.DeleteEmployee)
}
