package factories

import (
	"github.com/ismael3s/go-01/repositories"
	"github.com/ismael3s/go-01/repositories/implementations"
)

type ProblemRepositoryType string

const (
	CSV       ProblemRepositoryType = "CSV"
	IN_MEMORY ProblemRepositoryType = "IN_MEMORY"
	POSTGRES  ProblemRepositoryType = "POSTGRES"
)

func NewProblemRepository(p ProblemRepositoryType) repositories.ProblemRepository {
	switch p {
	case IN_MEMORY:
		return implementations.NewInMemoryProblemRepository()
	case POSTGRES:
		return implementations.NewPostgresProblemRepository()
	default:
		filename := "./assets/sum.csv"
		return implementations.NewCSVProblemRepository(filename)
	}
}
