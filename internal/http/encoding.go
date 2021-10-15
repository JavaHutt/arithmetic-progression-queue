package http

import (
	"encoding/json"
	"net/http"
)

type encoder struct{}

func newEncoder() *encoder {
	return &encoder{}
}

// errorResponse will encapsulate errors to be transferred over http.
type errorResponse struct {
	// return the reason of error
	Message string      `json:"message"`
	Trace   interface{} `json:"trace,omitempty"`
	Trail   interface{} `json:"trail,omitempty"`
}

func (e *encoder) JSONResponse(w http.ResponseWriter, response interface{}) {
	res, err := json.Marshal(response)
	if err != nil {
		e.Error(w, err, http.StatusInternalServerError)
		return
	}
	e.Response(w, res)
}

func (e *encoder) Response(w http.ResponseWriter, response []byte) {
	e.StatusResponse(w, response, http.StatusOK)
}

func (e *encoder) StatusOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (e *encoder) StatusResponse(w http.ResponseWriter, response []byte, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func (e *encoder) Error(w http.ResponseWriter, err error, status int) {
	resp := errorResponse{
		Message: err.Error(),
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
