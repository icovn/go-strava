package handler

import (
	"fmt"
	"github.com/icovn/go-strava/client"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[LoginHandler]")
	query := r.URL.Query()
	code := query.Get("activity")
	if code == "" {
		//redirect to login
	}
	w.WriteHeader(http.StatusOK)
	var stravaClient = client.StravaClient{Token: "xxx"}
	fmt.Fprintf(w, "Activities: %v\n", stravaClient.GetAthlete)
}
