package application

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zozoee27/cookbook/backend/account/manager"
	"github.com/zozoee27/cookbook/backend/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router         *mux.Router
	DatabaseClient *mongo.Client

	AccountManager *accountManager.AccountManager
}

func (a *App) Initialize(databaseName string) {
	a.Router = mux.NewRouter()
	connection := database.InitializeConnection(databaseName)
	a.DatabaseClient = connection.DatabaseClient

	a.AccountManager = accountManager.Initialize(connection.Database)
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
	database.Disconnect(a.DatabaseClient)
}

func (a *App) initializeRoutes() {
	log.Print("Initializing routes")
	a.Router.HandleFunc("/account/register", a.AccountManager.RegisterAccount).Methods("POST")
	a.Router.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Recieved request for /")
}
