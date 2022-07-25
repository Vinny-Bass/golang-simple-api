package main

import (
	"fmt"
	"simple-api/internal/comment"
	"simple-api/internal/db"
	transportHttp "simple-api/internal/transport/http"
)

func Run() error {
	provider, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := provider.MigrateDB(); err != nil {
		fmt.Println("Failed to run migrations")
		return err
	}

	cmtService := comment.NewService(provider)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	fmt.Println("Starting up our application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("There is an error in our application startup")
	}
}
