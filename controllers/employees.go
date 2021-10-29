package controllers

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Employees struct {
	ID          string
	Name        string
	Age         int
	Address     string
	Gender      string
	Department  string
	PhoneNumber string
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateEmployeeTable(db *bun.DB, ctx context.Context) string {
	_, err := db.NewCreateTable().
		Model((*Employees)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		return err.Error()
	} else {
		return "Table Created"
	}
}
