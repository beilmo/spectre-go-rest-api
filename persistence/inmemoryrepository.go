package persistence

// InMemoryRepository -
type InMemoryRepository struct {
}

// NewInMemoryRepository - constructor.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (repository InMemoryRepository) Foo() {

}
