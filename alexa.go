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
		cb := &RequestBody{}
		err := json.Unmarshal([]byte(body), cb)
		if err != nil {
			// TODO, add some logging here
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Get Traverser
		newTraverser := false
		traverser, err := store.FetchTraverser(cb.Session.SessionID)
		if err != nil {
			traverser, _ = store.CreateTraverser(cb.Session.SessionID)
			traverser.SetCurrentState("start")
			newTraverser = true
		}

		// Create Emitter
		emitter := &AlexaEmitter{
			ResponseWriter: w,
		}

		// Get Current State
		currentState := stateMachine[traverser.CurrentState()](emitter, traverser)
		if newTraverser {
			currentState.EntryAction()
		}

		// TODO plug in distiller
		// IntentDistiller <---
		distillerValue := ""

		// Transition
		newState := currentState.Transition(distillerValue)
		err = newState.EntryAction()
		if err == nil {
			traverser.SetCurrentState(newState.Slug)
		}

		// Write Body
		err = emitter.Flush()
		if err != nil {
			// TODO, add some logging here
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("response")
		// w.WriteHeader(http.StatusOK)
	}
}
