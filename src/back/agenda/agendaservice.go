package agenda

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
	"google.golang.org/api/calendar/v3"
	"github.com/Zenika/marcel-plugin-calendar/src/back/auth"
	"time"
	"os"
)

var calendarService *calendar.Service

func GetNextEvents(w http.ResponseWriter, r *http.Request) {

	var err error = nil

	Google_API_client := auth.RequireGoogleClient()
	calendarService, err = calendar.New(Google_API_client)

	if err != nil {
		log.Printf(err.Error() + "\n Requesting a new Client ID.")
		http.Redirect(w, r, "/api/v1/GoogleLogin", http.StatusTemporaryRedirect)
		return
	}

	var agendaId = os.Getenv("MARCEL_AGENDA_ID")

	//we want 'nbRequestedEvents' next events from today
	vars := mux.Vars(r)
	e := vars["nbEvents"]
	nbRequestedEvents, _ := strconv.Atoi(e)

	if nbRequestedEvents > 0 {
		var startTime time.Time = time.Now() //today

		calendarEvents, err := calendarService.Events.List(agendaId).
			SingleEvents(true).
			TimeMin(startTime.Format(time.RFC3339)).
			MaxResults(int64(nbRequestedEvents)).
			OrderBy("startTime").
			Do()

		if err != nil {
			w.Write([]byte(err.Error()))
			log.Fatal(err)
		}

		//Example of parsing an event
		/*if len(calendarEvents.Items) > 0 {
			for _, i := range calendarEvents.Items {
				var when string
				// If the DateTime is an empty string the Event is an all-day Event.
				// So only Date is available.
				if i.Start.DateTime != "" {
					when = i.Start.DateTime
				} else {
					when = i.Start.Date
				}
				fmt.Printf("%s (%s)\n", i.Summary, when)
			}
		} else {
			fmt.Printf("No upcoming events found.\n")
		}*/

		j, err := json.Marshal(calendarEvents)
		if err == nil {
			w.Write([]byte(j))
		}
	}
}