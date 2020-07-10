package client

import (
	"encoding/json"
	"github.com/icovn/go-strava/model"
	"io/ioutil"
	"log"
	"net/http"
)

const apiRoot = "https://www.strava.com/api/v3"
const userAgent = "PostmanRuntime/7.26.1"
var client = &http.Client{}

type StravaClient struct {
	Token string
}

func (a *StravaClient) GetAthlete() model.Athlete {
	log.Printf("GetAthlete")
	req, err := http.NewRequest("GET", apiRoot + "/athlete", nil)
	req.Header.Add("Authorization", "Bearer " + a.Token)
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		return model.Athlete{}
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		athlete := model.Athlete{}
		json.Unmarshal(data, athlete)
		return athlete
	}
}

func (a *StravaClient) GetAthleteActivities() []model.Activity {
	log.Printf("GetAthleteActivities")
	req, err := http.NewRequest("GET", apiRoot + "/athlete/activities?before=1&after=2&page=1&per_page=30", nil)
	req.Header.Add("Authorization", "Bearer " + a.Token)
	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("GetAthleteActivities: " + err.Error())
		return nil
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var activities []model.Activity
		json.Unmarshal(data, activities)
		return activities
	}
}
