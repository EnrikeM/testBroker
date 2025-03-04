package httpsrv

import (
	"TestBroker/internal/queue"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (a *API) handleMsg(w http.ResponseWriter, r *http.Request) {
	queueName := chi.URLParam(r, "queue")
	if queueName == "" {
		writeResponse(w, http.StatusBadRequest, "queue cannot be nil")

		return
	}

	reqBytes := json.NewDecoder(r.Body)
	var msg queue.Message
	if err := reqBytes.Decode(&msg); err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Errorf("unmarshal err: %w", err))

		return
	}

	err := a.queues.SetMessage(msg, queueName)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, fmt.Errorf("setMessage err: %w", err))

		return
	}

	writeResponse(w, http.StatusOK, "")
}

func (a *API) handleSub(w http.ResponseWriter, r *http.Request) {
	queue := chi.URLParam(r, "queue")
	if queue == "" {
		writeResponse(w, http.StatusBadRequest, "queue cannot be nil")

		return
	}

}
