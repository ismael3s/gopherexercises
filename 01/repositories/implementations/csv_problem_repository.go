package implementations

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/ismael3s/go-01/entities"
	"github.com/ismael3s/go-01/repositories"
	"github.com/ismael3s/go-01/util"
)

var _ repositories.ProblemRepository = new(CSVProblemRepository)

type CSVProblemRepository struct {
	Filename string
}

func (r CSVProblemRepository) Read() []entities.Problem {
	f, err := os.Open(r.Filename)

	eP := util.NewErrorParam(err, fmt.Sprintf("Failed to open file: %s", r.Filename), util.Fatal)
	util.HandleError(eP)

	defer f.Close()

	lines := readCsv(f)

	problems := parseLines(lines)

	return problems
}

func NewCSVProblemRepository(filename string) *CSVProblemRepository {
	return &CSVProblemRepository{Filename: filename}
}

func readCsv(file *os.File) (lines [][]string) {
	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()

	eP := util.NewErrorParam(err, fmt.Sprintf("Failed to parse provide CSV.\nError: %s", err), util.Fatal)
	util.HandleError(eP)

	return lines
}

func parseLines(lines [][]string) (ret []entities.Problem) {
	ret = make([]entities.Problem, len(lines))

	for i, line := range lines {
		ret[i] = entities.NewProblem(line[0], strings.TrimSpace(line[1]))
	}

	return ret
}
