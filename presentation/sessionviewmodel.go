package presentation

import (
	"errors"

	"github.com/beilmo/spectre-go-rest-api/domain"
	"github.com/beilmo/spectre-go-rest-api/presentation/dto"
	"github.com/beilmo/spectre-go-rest-api/presentation/serializers"
)

// SessionViewModel type.
type SessionViewModel struct {
	useCaseFactory *domain.UseCaseFactory
}

// NewSessionViewModel - constructor that injects the useCaseFactory.
func NewSessionViewModel(useCaseFactory *domain.UseCaseFactory) *SessionViewModel {
	return &SessionViewModel{
		useCaseFactory: useCaseFactory,
	}
}

// GetAllSessions - returns all the sessions.
func (vm *SessionViewModel) GetAllSessions() (dto.SessionList, error) {
	// 1. Create corresponding useCase.
	uc := vm.useCaseFactory.MakeAllSessions()

	// 2. Execute the useCase.
	ch := make(chan domain.Result)
	go uc.Execute(ch)

	// 3. Get the result out of the channel.
	r := <-ch

	// 4. Check for errors.
	if r.Error != nil {
		return dto.SessionList{}, r.Error
	}

	// 5. Cast the result from the empty interface.
	t, ok := r.Value.([]domain.Session)
	if ok == false {
		return dto.SessionList{}, errors.New("Unexpected type. Was looking for []domain.Session")
	}

	// 6. Encode domain model into dto model.

	n := len(t)

	var serializedList []*dto.Session
	for i := 0; i < n; i++ {
		session := serializers.EncodeSession(t[i])
		serializedList = append(serializedList, &session)
	}

	list := dto.SessionList{
		Elements: serializedList,
	}

	// 7. Return the encoded model
	return list, nil
}

// GetSession - returns the session identified by the specified id.
// If not found, an error will be returned.
func (vm *SessionViewModel) GetSession(id int64) (dto.Session, error) {
	// 1. Create corresponding useCase.
	uc := vm.useCaseFactory.MakeSessionByID(id)

	// 2. Execute the useCase.
	ch := make(chan domain.Result)
	go uc.Execute(ch)

	// 3. Get the result out of the channel.
	r := <-ch

	// 4. Check for errors.
	if r.Error != nil {
		return dto.Session{}, r.Error
	}

	// 5. Cast the result from the empty interface.
	t, ok := r.Value.(domain.Session)
	if ok == false {
		return dto.Session{}, errors.New("Unexpected type. Was looking for domain.Session")
	}

	// 6. Encode domain model into dto model.
	session := serializers.EncodeSession(t)

	// 7. Return the encoded model
	return session, nil
}
