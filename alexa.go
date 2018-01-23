package fsmalexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/TV4/graceful"
	"github.com/go-carrot/fsm"
	"github.com/julienschmidt/httprouter"
)

func Start(stateMachine fsm.StateMachine, store fsm.Store) {
	graceful.LogListenAndServe(
		&http.Server{
			Addr:    ":" + os.Getenv("PORT"),
			Handler: buildRouter(stateMachine, store),
		},
	)
}

func buildRouter(stateMachine fsm.StateMachine, store fsm.Store) *httprouter.Router {
	// Router
	router := &httprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
	}
	router.HandlerFunc(http.MethodPost, "/alexa", getAlexaWebhook(stateMachine, store))
	return router
}

//
func getAlexaWebhook(stateMachine fsm.StateMachine, store fsm.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get Body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		// Parse body into struct
		cb := &MessageReceivedRequestBody{}
		json.Unmarshal([]byte(body), cb)

		fmt.Println(cb.Request.Intent)
		w.WriteHeader(http.StatusOK)
	}
}
