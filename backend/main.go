package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"

    "github.com/zozoee27/cookbook/backend/database"
    "github.com/zozoee27/cookbook/backend/account/manager"
)

func main() {
    router := mux.NewRouter()

    initDatabase()
    initAccountOps(router)

    log.Print("Cookbook server is running and listening on port: 8080")
    http.ListenAndServe(":8080" , router)
}

func initAccountOps(router *mux.Router) {
    router.HandleFunc("/account/register", accountManager.RegisterAccount).Methods("POST");

}

func initDatabase() {
    database.InitializeConnection()
}


