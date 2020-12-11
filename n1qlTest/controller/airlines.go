package controller

import (
	"encoding/json"
	"net/http"

	"n1qlServerTest/model/travelsample"
	. "n1qlServerTest/util"
)

func AllAirlinesController(w http.ResponseWriter, r *http.Request) {
	airlineData, err := travelsample.GetAllAirlines()
	if err != nil {
		LoggerError.Println("Failed to get all airlines")
		return
	}
	json.NewEncoder(w).Encode(airlineData)
}
