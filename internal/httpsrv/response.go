package httpsrv

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, status int, r any) {
	b, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(b); err != nil {
		return
	}
}
