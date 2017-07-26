package main

import (
	"github.com/Zenika/marcel-plugin-calendar/src/back/agenda"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const SERVICE_URL = "/agenda/incoming/"
const SERVICE_ATTR = "{nbEvents:[0-9]*}"

func main() {
	logFile := "/tmp/marcel.plugins.official.agenda.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	defer f.Close()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	SetRoutes(r)
	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}

func SetRoutes(r *mux.Router) {
	r.HandleFunc(SERVICE_URL + SERVICE_ATTR, agenda.GetNextEvents).Methods("GET")
}