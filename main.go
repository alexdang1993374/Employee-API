package main

import (
	"context"

	"github.com/alexdang1993374/employee-api/config"
	"github.com/alexdang1993374/employee-api/controllers"
)

func main() {
	db := config.Connect()
	ctx := context.Background()

	controllers.CreateEmployeeTable(db, ctx)
}
