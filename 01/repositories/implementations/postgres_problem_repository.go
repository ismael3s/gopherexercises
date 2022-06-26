package implementations

import (
	"database/sql"
	"log"

	"github.com/ismael3s/go-01/entities"
	"github.com/ismael3s/go-01/repositories"
	"github.com/ismael3s/go-01/util"
	_ "github.com/lib/pq"
)

var _ repositories.ProblemRepository = new(PostgresProblemRepository)

type PostgresProblemRepository struct {
	db *sql.DB
}

func (r PostgresProblemRepository) Read() []entities.Problem {
	var problems []entities.Problem
	rows, err := r.db.Query(`SELECT question, answer FROM "problem"`)

	util.HandleError(&util.ErrorParam{Err: err, Message: "Failed on Query", Flow: util.Fatal})

	for rows.Next() {
		var problem entities.Problem
		if err := rows.Scan(&problem.Question, &problem.Answer); err != nil {
			log.Fatal(err.Error())
		}

		problems = append(problems, problem)
	}

	defer r.db.Close()
	defer rows.Close()

	return problems
}

func NewPostgresProblemRepository() *PostgresProblemRepository {
	db := connectToDb()
	return &PostgresProblemRepository{db: db}
}

func connectToDb() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5433 user=root password=root dbname=aroot sslmode=disable")

	util.HandleError(util.NewErrorParam(err, "Failed to Open database connection", util.Fatal))

	return db
}
