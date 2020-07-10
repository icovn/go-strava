package handler

import (
	"fmt"
	"github.com/icovn/go-strava/client"
	"log"
	"net/http"
)

func ActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ActivitiesHandler]")
	query := r.URL.Query()
	activity := query.Get("activity")
	if activity == "" {
		activity = "N/A"
	}
	w.WriteHeader(http.StatusOK)
	var stravaClient = client.StravaClient{Token: "xxx"}
	fmt.Fprintf(w, "Activities: %v\n", stravaClient.GetAthlete)
}
