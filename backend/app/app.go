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
	router             *mux.Router
	databaseConnection *database.Connection
	userService        *service.User
}

func (a *App) Initialize(databaseName string) {
	a.router = mux.NewRouter()

	a.databaseConnection = &database.Connection{}
	a.databaseConnection.InitializeConnection(databaseName)

	a.initializeRoutes()
}

func (a *App) Run(address string) {
	log.Print("Cookbook server is running and listening on port: ", address)
	err := http.ListenAndServeTLS(address, "/etc/ssl/localhost-certs/localhost.crt", "/etc/ssl/localhost-certs/localhost.key", a.router)

	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}

func (a *App) StopApplication() {
	a.databaseConnection.Disconnect()
}

func (a *App) initializeRoutes() {
	log.Print("Initializing routes")
	handlers.MakeUserHandlers(a.router, a.userService)
}

func (a *App) createUserServices() {
	db := database.CreateUserMongoDb(a.databaseConnection.Database.Collection("users"))
	a.userService = service.CreateUserService(db)

}
