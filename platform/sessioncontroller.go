package platform

import (
	"net/http"
	"strconv"

	"github.com/beilmo/spectre-go-rest-api/presentation"
	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/mux"
)

// SessionController type.
type SessionController struct {
	viewModel *presentation.SessionViewModel
}

// NewSessionController - constructor that injects the viewModel.
func NewSessionController(viewModel *presentation.SessionViewModel) *SessionController {
	return &SessionController{
		viewModel: viewModel,
	}
}

// Routes - Routable interface implementation providing routes to this controller.
func (controller SessionController) Routes() Routes {
	return Routes{
		Route{"ListAllSessions", "GET", "/api/sessions", controller.List},
		Route{"FindSessionByID", "GET", "/api/sessions/{id:[0-9]+}", controller.FindByID},
	}
}

// List -
func (controller *SessionController) List(w http.ResponseWriter, r *http.Request) {
	s, _ := controller.viewModel.GetAllSessions()

	m := jsonpb.Marshaler{}

	w.Header().Add("Content-Type", "application/json")
	m.Marshal(w, &s)
}

// FindByID -
func (controller *SessionController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	s, _ := controller.viewModel.GetSession(id)

	m := jsonpb.Marshaler{}

	w.Header().Add("Content-Type", "application/json")
	m.Marshal(w, &s)
}
