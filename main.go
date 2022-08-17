package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AfandyW/shopping-cart/app"
	"github.com/AfandyW/shopping-cart/controller"
	"github.com/AfandyW/shopping-cart/repository"
	"github.com/AfandyW/shopping-cart/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Fail to load env", err)
		return
	}

	host := os.Getenv("HOST")
	dbport := os.Getenv("PORT")
	serverPort := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")

	db, err := app.NewDB(dbport, host, dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println("Check setup DB", err)
		return
	}

	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	r := app.NewRouter(*controller)

	server := http.Server{
		Addr:    host + ":" + serverPort,
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to run server", err)
		return
	}

}
