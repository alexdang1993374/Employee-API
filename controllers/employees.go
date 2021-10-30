package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"github.com/uptrace/bun"
)

var dbConnect *bun.DB

func InitiateDB(db *bun.DB) {
	dbConnect = db
}

type Employees struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
	Department  string    `json:"department"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateEmployeeTable(ctx context.Context) string {
	_, err := dbConnect.NewCreateTable().
		Model((*Employees)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		return err.Error()
	} else {
		return "Table Created"
	}
}

func CreateEmployee(c *gin.Context) {
	employee := Employees{}

	c.BindJSON(&employee)

	// id := guuid.New().String()
	// name := employee.Name
	// age := employee.Age
	// address := employee.Address
	// gender := employee.Gender
	// department := employee.Department
	// phone_number := employee.PhoneNumber

	employee = Employees{
		ID:          guuid.New().String(),
		Name:        employee.Name,
		Age:         employee.Age,
		Address:     employee.Address,
		Gender:      employee.Gender,
		Department:  employee.Department,
		PhoneNumber: employee.PhoneNumber,
	}

	_, err := dbConnect.NewInsert().Model(&employee).Exec(c)
	if err != nil {
		log.Printf("Error while inserting new employee into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Employee created Successfully",
		})
		return
	}
}
