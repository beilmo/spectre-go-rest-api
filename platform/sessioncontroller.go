package platform

import (
	"net/http"

	"github.com/beilmo/spectre-go-rest-api/presentation"
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
func (controller *SessionController) List(responseWriter http.ResponseWriter, r *http.Request) {
}

// FindByID -
func (controller *SessionController) FindByID(responseWriter http.ResponseWriter, r *http.Request) {
}
