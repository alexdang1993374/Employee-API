package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

var dbConnect *bun.DB

func InitiateDB(db *bun.DB) {
	dbConnect = db
}

type Employees struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
	Department  string    `json:"department"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateEmployeeTable(ctx context.Context) error {
	_, err := dbConnect.NewCreateTable().
		Model((*Employees)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error while creating login table, Reason: %v\n", err)
		return err
	} else {
		log.Printf("Login table created")
		return nil
	}
}

func CreateEmployee(c *gin.Context) {

	employee := Employees{}

	c.BindJSON(&employee)

	employee = Employees{
		ID:          employee.ID,
		FirstName:   employee.FirstName,
		LastName:    employee.LastName,
		Age:         employee.Age,
		Address:     employee.Address,
		Gender:      employee.Gender,
		Department:  employee.Department,
		PhoneNumber: employee.PhoneNumber,
	}

	exists, findErr := dbConnect.NewSelect().Model((*Employees)(nil)).Where("id = ?", employee.ID).Exists(c)
	if findErr != nil {
		panic(findErr)
	}
	if exists {
		log.Printf("Error while inserting new employee into db, Reason: Employee already exists")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error while inserting new employee into db, Reason: Employee already exists",
		})
		return
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
			"data":    employee,
		})
		return
	}
}

func UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("employeeId")
	newEmployee := Employees{}
	oldEmployee := Employees{}

	err := dbConnect.NewSelect().Model((*Employees)(nil)).Where("id = ?", employeeID).Scan(c, &oldEmployee)
	if err != nil {
		log.Printf("Error while getting an employee, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Employee not found",
		})
		return
	}

	c.BindJSON(&newEmployee)

	id := oldEmployee.ID
	var firstName string
	var lastName string
	var age int
	var address string
	var gender string
	var department string
	var phoneNumber string

	if newEmployee.FirstName != "" {
		firstName = newEmployee.FirstName
	} else {
		firstName = oldEmployee.FirstName
	}

	if newEmployee.LastName != "" {
		lastName = newEmployee.LastName
	} else {
		lastName = oldEmployee.LastName
	}

	if newEmployee.Age != 0 {
		age = newEmployee.Age
	} else {
		age = oldEmployee.Age
	}

	if newEmployee.Address != "" {
		address = newEmployee.Address
	} else {
		address = oldEmployee.Address
	}

	if newEmployee.Gender != "" {
		gender = newEmployee.Gender
	} else {
		gender = oldEmployee.Gender
	}

	if newEmployee.Department != "" {
		department = newEmployee.Department
	} else {
		department = oldEmployee.Department
	}

	if newEmployee.PhoneNumber != "" {
		phoneNumber = newEmployee.PhoneNumber
	} else {
		phoneNumber = oldEmployee.PhoneNumber
	}

	newEmployee = Employees{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Age:         age,
		Address:     address,
		Gender:      gender,
		Department:  department,
		PhoneNumber: phoneNumber,
	}

	_, updateError := dbConnect.NewUpdate().
		Model(&newEmployee).
		Column("first_name").
		Column("last_name").
		Column("age").
		Column("address").
		Column("gender").
		Column("department").
		Column("phone_number").
		Where("id = ?", employeeID).
		Exec(c)

	if updateError != nil {
		log.Printf("Error, Reason: %v\n", updateError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Employee Edited Successfully",
			"data":    &newEmployee,
		})
		return
	}
}

func DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("employeeId")

	oldEmployee := Employees{}

	err := dbConnect.NewSelect().Model((*Employees)(nil)).Where("id = ?", employeeID).Scan(c, &oldEmployee)
	if err != nil {
		log.Printf("Error while getting an employee, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Employee not found",
		})
		return
	}

	_, deleteError := dbConnect.NewDelete().
		Model((*Employees)(nil)).
		Where("id = ?", employeeID).
		Exec(c)

	if deleteError != nil {
		log.Printf("Error while deleting a single employee, Reason: %v\n", deleteError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employee deleted successfully",
		})
		return
	}
}

func GetAllEmployees(c *gin.Context) {
	var employees []Employees

	_, err := dbConnect.NewSelect().Model(&employees).Order("first_name ASC").Limit(999).ScanAndCount(c)
	if err != nil {
		log.Printf("Error while getting all employees, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Employees",
		"count":   len(employees),
		"data":    employees,
	})
	return
}
