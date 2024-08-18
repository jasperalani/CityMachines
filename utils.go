package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateErrorResponse(w http.ResponseWriter, queryColumn string) bool {

	results := DB.QueryRow("SELECT " + queryColumn + " FROM errors;")

	response := &Response{
		Errno: 1,
	}

	err = results.Scan(&response.Error)
	HandleError(err)

	data, err := json.Marshal(response)
	HandleError(err)

	_, err = w.Write(data)
	HandleError(err)

	return true
	//not sure why i am returning true here

}

func CreateResponse(w http.ResponseWriter, response string) {

	responseObj := &Response{
		Response: response,
		Errno:    0,
	}

	data, err := json.Marshal(responseObj)
	HandleError(err)

	_, err = w.Write(data)
	HandleError(err)

}

func StringEvaluator(subject bool, outcome1 string, outcome2 string) string {
	if subject {
		return outcome1
	}
	return outcome2
}

//func FunctionEvaluator (subject bool, outcome1 func(Type reflect.Type))

func SelectID(id string) float32 {

	/*

		Error codes:
		-11 sql: no rows in result set

	*/

	var (
		itemExists bool
		queryID    = "SELECT id FROM items WHERE id = " + id
		item       Item
		err        error
	)

	err = DB.Get(&item, queryID)

	if err == sql.ErrNoRows {
		return 0
	} else {
		HandleError(err)
	}

	itemExists = item.ID != 0

	if !itemExists {
		return 0
	}

	return float32(item.ID)

}

func AnyItems() bool {

	var (
		items    []Item
		queryAll = "SELECT * FROM `shopping-list`.items;"
	)

	err = DB.Select(&items, queryAll)
	return err != sql.ErrNoRows

}
