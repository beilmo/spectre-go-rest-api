package router

import (
	"net/http"
	"strconv"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/mux"

	"github.com/beilmo/spectre-go-rest-api/application"
	"github.com/beilmo/spectre-go-rest-api/application/repository"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/dto"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/presenters"
	"github.com/beilmo/spectre-go-rest-api/interfaces/controllers"
)

// SessionsRequestHandler type.
type SessionsRequestHandler struct {
	Logger  application.Logger
	Storage repository.SessionRepository
}

// FindAllSessions -
func (h SessionsRequestHandler) FindAllSessions(w http.ResponseWriter, r *http.Request) {
	controller := controllers.NewSessionController(h.Storage, h.Logger)
	presenter := presenters.SessionsListPresenter{
		Handler: func(sessions dto.SessionList, err error) {
			m := jsonpb.Marshaler{}
			out, err := m.MarshalToString(&sessions)
			h.Logger.Log(out)

			w.Header().Add("Content-Type", "application/json")
			m.Marshal(w, &sessions)
		},
	}

	controller.ListAllSessions(presenter)
}

// FindSessionByID -
func (h SessionsRequestHandler) FindSessionByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	controller := controllers.NewSessionController(h.Storage, h.Logger)
	presenter := presenters.SessionPresenter{
		Handler: func(session dto.Session, err error) {
			m := jsonpb.Marshaler{}
			out, err := m.MarshalToString(&session)
			h.Logger.Log(out)

			w.Header().Add("Content-Type", "application/json")
			m.Marshal(w, &session)
		},
	}

	controller.ListSessionByID(id, presenter)
}
