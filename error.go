package main

import (
	"log"
	"net/http"
)

var (
	err      error
	errorMsg = "Unable to process request"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HTTPNotFound(w http.ResponseWriter, _ *http.Request) {
	CreateResponse(w, errorMsg, "502", "HTTP Not Found")
}

// Create IDNotFound error response
//func IDNotFound(w http.ResponseWriter, _ *http.Request) {
//	CreateResponse(w, "Unable to process request", "301", "HTTP Not Found")
//}

// Create NoDataProvided error response
//func NoDataProvided(w http.ResponseWriter, _ *http.Request) {
//	CreateErrorResponse(w, "err_nodataprovided")
//}

func NoItems(w http.ResponseWriter, _ *http.Request) {
	CreateResponse(w, errorMsg, "1", "Search returned 0 results.")
}
