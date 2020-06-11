package util

import (
	"encoding/json"
	"net/http"
)

// MarshalAndWriteResponse marshals the interface to JSON and then writes it to the ResponseWriter
func MarshalAndWriteResponse(w http.ResponseWriter, r *http.Request, i interface{}) {
	json, err := json.Marshal(i)
	if err != nil {
		ReturnError(w, r, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(json))
}

// ReturnError writes an error back to the client
func ReturnError(w http.ResponseWriter, r *http.Request, err error, code int) {
	http.Error(w, err.Error(), code)
}
