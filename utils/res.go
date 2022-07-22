package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(rw http.ResponseWriter, p interface{}, status int) {
	changeToByte, err := json.Marshal(p)
	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.WriteHeader(status)
	rw.Write([]byte(changeToByte))
}
