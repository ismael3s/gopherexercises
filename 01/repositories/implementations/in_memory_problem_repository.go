package implementations

import (
	"github.com/ismael3s/go-01/entities"
	"github.com/ismael3s/go-01/repositories"
)

var _ repositories.ProblemRepository = new(InMemoryProblemRepository)

type InMemoryProblemRepository struct{}

func (r InMemoryProblemRepository) Read() []entities.Problem {
	problems := []entities.Problem{{Question: "2 + 1", Answer: "3"}, {Question: "2 + 2", Answer: "4"}}

	return problems
}

func NewInMemoryProblemRepository() *InMemoryProblemRepository {
	return &InMemoryProblemRepository{}
}
