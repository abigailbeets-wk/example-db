package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type HttpErrorResponse struct {
	ErrorDetails HttpErrorDetails `json:"error"`
}

type HttpErrorDetails struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	httpErr := HttpErrorResponse{
		ErrorDetails: HttpErrorDetails{
			Message: message,
			Code:    strconv.Itoa(code),
		},
	}

	RespondWithJSON(w, code, httpErr)
}

func RespondWithStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
