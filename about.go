package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func About(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Response: `CityMachines API at ` + time.Now().Format("2006-01-02 15:04:05"),
		Errno:    0,
		Error:    `None`,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		HandleError(err)
	}
}
