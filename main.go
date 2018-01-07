package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/atiernan/smartHomeSamsungTVCommon"
)

var currentResponse = smartHomeSamsungTVCommon.DeviceEndpointResponse{}

func googleAssistantEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(p)
}

func deviceEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentResponse)
	currentResponse = smartHomeSamsungTVCommon.DeviceEndpointResponse{
		TVSwitchedOn:  false,
		TVSwitchedOff: false,
		VolumeUp:      0,
		VolumeDown:    0,
		VolumeMute:    false,
		Pause:         false,
		Play:          false,
		OK:            false,
	}
}

func uiEndpoint(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query.Get("TVOn")) != 0 {
		currentResponse.TVSwitchedOn = true
	}
	if len(query.Get("TVOff")) != 0 {
		currentResponse.TVSwitchedOff = true
	}
	if val := query.Get("VolumeUp"); len(val) != 0 {
		i, err := strconv.Atoi(val)
		if err == nil {
			currentResponse.VolumeUp += i
		}
	}
	if val := query.Get("VolumeDown"); len(val) != 0 {
		i, err := strconv.Atoi(val)
		if err == nil {
			currentResponse.VolumeDown += i
		}
	}
	if len(query.Get("VolumeMute")) != 0 {
		currentResponse.VolumeMute = true
	}
	if len(query.Get("Pause")) != 0 {
		currentResponse.Pause = true
	}
	if len(query.Get("Play")) != 0 {
		currentResponse.Play = true
	}
	if len(query.Get("OK")) != 0 {
		currentResponse.OK = true
	}

	fmt.Fprintln(w, "")
}

func main() {
	http.HandleFunc("/google-assistant/endpoint", googleAssistantEndpoint)
	http.HandleFunc("/device/endpoint", deviceEndpoint)
	http.HandleFunc("/", uiEndpoint)
	http.ListenAndServe(":8081", nil)
}
