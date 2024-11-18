package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func sendResponse(w http.ResponseWriter, httpStatus int, response Response) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	responseBody, _ := json.Marshal(response)
	w.Write([]byte(responseBody))
}
