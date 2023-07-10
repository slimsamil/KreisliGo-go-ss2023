package main

import (
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/handler"

)

func init() {
	//ensure that logger is initialized before connecting to DB
	defer db.Init()
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}


func main() {
	fmt.Println("Starting Kreisligo API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/association", handler.CreateAssociation).Methods("POST")
	router.HandleFunc("/associations", handler.GetAssociations).Methods("GET")
	router.HandleFunc("/association/{id}", handler.GetAssociation).Methods("GET")
	router.HandleFunc("/association/{id}", handler.UpdateAssociation).Methods("PUT")
	router.HandleFunc("/association/{id}", handler.DeleteAssociation).Methods("DELETE")
	router.HandleFunc("/association/{id}/leagues", handler.GetLeagues).Methods("GET")
	router.HandleFunc("/association/{id}/league", handler.CreateLeague).Methods("POST")
	router.HandleFunc("/association/{id}/league/{id}", handler.GetLeague).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}", handler.UpdateLeague).Methods("PUT")
	router.HandleFunc("/association/{id}/league/{id}", handler.DeleteLeague).Methods("DELETE")
	router.HandleFunc("/association/{id}/league/{id}/games", handler.GetGames).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/game", handler.CreateGame).Methods("POST")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}", handler.GetGame).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}", handler.UpdateGame).Methods("PUT")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}", handler.DeleteGame).Methods("DELETE")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}/events", handler.GetEvents).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/association/{id}/league/{id}/game/{id}/event/{id}", handler.GetEvent).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/teams", handler.GetTeams).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/team", handler.CreateTeam).Methods("POST")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}", handler.GetTeam).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}", handler.UpdateTeam).Methods("PUT")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}", handler.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}/players", handler.GetPlayers).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}/player", handler.CreatePlayer).Methods("POST")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}/player/{id}", handler.GetPlayer).Methods("GET")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}/player/{id}", handler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/association/{id}/league/{id}/team/{id}/player/{id}", handler.DeletePlayer).Methods("DELETE")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}