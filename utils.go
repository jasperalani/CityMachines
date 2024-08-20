package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func CreateResponse(w http.ResponseWriter, response string, error ...string) {
	errorNumber := 0
	errorString := ""
	if len(error) > 1 {
		errorNumber, err = strconv.Atoi(error[0])
		HandleError(err)
		errorString = error[1]
	}

	responseObj := &Response{
		Response: response,
		Errno:    errorNumber,
		Error:    errorString,
	}

	data, err := json.Marshal(responseObj)
	HandleError(err)

	_, err = w.Write(data)
	HandleError(err)
}

func getNodeByID(id string) (*Node, error) {
	var (
		blank = `SELECT * FROM machines WHERE id = 1`
		query = strings.Replace(blank, `1`, id, 1)
		node  Node
		err   error
	)

	err = DB.Get(&node, query)

	if err != nil {
		return nil, err
	}

	return &node, nil
}
