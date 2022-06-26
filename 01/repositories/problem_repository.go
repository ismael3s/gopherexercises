package repositories

import "github.com/ismael3s/go-01/entities"

type ProblemRepository interface {
	Read() []entities.Problem
}
