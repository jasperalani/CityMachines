package main

import (
	"log"
	"net/http"
)

var err error

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Create HTTPNotFound error response
// Need both w and r to be a valid http.HandlerFunc type
func HTTPNotFound(w http.ResponseWriter, _ *http.Request) {
	CreateErrorResponse(w, "err_httpnotfound")
}

// Create IDNotFound error response
func IDNotFound(w http.ResponseWriter, _ *http.Request) {
	CreateErrorResponse(w, "err_idnotfound")
}

// Create NoDataProvided error response
//func NoDataProvided(w http.ResponseWriter, _ *http.Request) {
//	CreateErrorResponse(w, "err_nodataprovided")
//}

// Create NoItems error response
func NoItems(w http.ResponseWriter, _ *http.Request) {
	CreateErrorResponse(w, "err_noitems")
}
