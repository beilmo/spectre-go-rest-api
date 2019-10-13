package presenters

import (
	"github.com/beilmo/spectre-go-rest-api/domain/model"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/dto"
	"github.com/beilmo/spectre-go-rest-api/interfaces/api/serializers"
)

// SessionsListPresenter type
type SessionsListPresenter struct {
	Handler func(dto.SessionList, error)
}

// PresentSessionsList - ListSessionsOutputBoundary interface implementation
func (p SessionsListPresenter) PresentSessionsList(sessions []model.Session, err error) {
	if err != nil {
		p.Handler(dto.SessionList{}, err)
		return
	}

	n := len(sessions)

	var serializedList []*dto.Session
	for i := 0; i < n; i++ {
		session := serializers.EncodeSession(sessions[i])
		serializedList = append(serializedList, &session)
	}

	list := dto.SessionList{
		Elements: serializedList,
	}
	p.Handler(list, nil)
}

type SessionPresenter struct {
	Handler func(dto.Session, error)
}

func (p SessionPresenter) PresentSession(session model.Session, err error) {
	if err != nil {
		p.Handler(dto.Session{}, err)
		return
	}

	p.Handler(serializers.EncodeSession(session), nil)
}
