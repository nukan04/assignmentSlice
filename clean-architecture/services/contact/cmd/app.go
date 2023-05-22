package main

import (
	"fmt"

	"nurda/pkg/store/postgres"
	"nurda/services/contact/internal/domain"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234567",
		DbName:   "Day1",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	alexey := domain.NewContact("Fedenko", "Alexey", "Demyanovich")
	nurda := domain.NewContact("Nurdaulet", "Kuatov", "Nurbolatovich")
	group1 := domain.NewGroup("Students")

	fmt.Println(alexey, nurda, group1)
}
