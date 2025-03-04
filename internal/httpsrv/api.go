package httpsrv

import (
	"TestBroker/internal/queue"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type API struct {
	queues queue.Queues
	router chi.Router
}

func NewAPI() *API {
	api := &API{
		router: chi.NewRouter(),
	}

	api.register(api.router)
	return api
}

func (a *API) SetQueues(newQueues queue.Queues) {
	a.queues = newQueues
}

func (a *API) Start() error {
	err := http.ListenAndServe("localhost:8080", a.router)
	if err != nil {
		return fmt.Errorf("error starting server %w", err)
	}

	return nil
}

func (a *API) register(r chi.Router) {
	r.Route("/v1/queues/{queue}", func(r chi.Router) {
		r.Post("/messages", a.handleMsg)
		r.Post("/subscriptions", a.handleSub)
	})
}

// POST ${address}/v1/queues/${queue_name}/messages
// POST ${address}/v1/queues/${queue_name}/subscriptions
