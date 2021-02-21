package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zozoee27/cookbook/backend/database"
	"github.com/zozoee27/cookbook/backend/handlers"
	"github.com/zozoee27/cookbook/backend/service"
)

type App struct {
	Router             *mux.Router
	DatabaseConnection *database.Connection
	UserService        *service.User
}

func (a *App) Initialize(databaseName string) {
	a.Router = mux.NewRouter()

	a.DatabaseConnection = &database.Connection{}
	a.DatabaseConnection.InitializeConnection(databaseName)

	a.createUserServices()

	a.initializeRoutes()
}

func (a *App) Run(address string) {
	log.Print("Cookbook server is running and listening on port: ", address)
	err := http.ListenAndServeTLS(address, "/etc/ssl/localhost-certs/localhost.crt", "/etc/ssl/localhost-certs/localhost.key", a.Router)

	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}

func (a *App) StopApplication() {
	a.DatabaseConnection.Disconnect()
}

func (a *App) initializeRoutes() {
	log.Print("Initializing routes")
	handlers.MakeUserHandlers(a.Router, a.UserService)
}

func (a *App) createUserServices() {
	db := database.CreateUserMongoDb(a.DatabaseConnection.Database.Collection("users"))
	a.UserService = service.CreateUserService(db)

}
