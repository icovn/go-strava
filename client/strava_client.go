package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const apiRoot = "https://www.strava.com/api/v3"
const userAgent = "PostmanRuntime/7.26.1"
var client = &http.Client{}

func GetAthlete(token string) {
	req, err := http.NewRequest("GET", apiRoot + "/athlete", nil)
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Printf(string(data))
	}

}

func GetAthleteActivities(token string) {
	req, err := http.NewRequest("GET", apiRoot + "/athlete/activities?before=1&after=2&page=1&per_page=30", nil)
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Printf(string(data))
	}
}
