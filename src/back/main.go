package main

import (
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"log"
	"github.com/Zenika/marcel-plugin-calendar/src/back/agenda"
)

func main() {

	logFile := "marcel.plugins.official.agenda.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	defer f.Close()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	s := r.PathPrefix("/plugins/official/").Subrouter()
	s.HandleFunc("/agenda/incoming/{nbEvents:[0-9]*}", agenda.GetNextEvents).Methods("GET")
	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}

