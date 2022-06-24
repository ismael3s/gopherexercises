package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ismael3s/go-01/util"
)

type problem struct {
	question, answer string
}

func main() {
	filename := "./assets/sum.csv"
	timeLimit := 1
	f, err := os.Open(filename)

	eP := util.NewErrorParam(err, fmt.Sprintf("Failed to open file: %s", filename), util.Fatal)
	util.HandleError(eP)

	defer f.Close()

	lines := readCsv(f)

	problems := parseLines(lines)

	startSolving(problems, timeLimit)

}

func startSolving(problems []problem, timeLimit int) (score int) {
	timer := time.NewTimer(time.Second * time.Duration(timeLimit))
	score = 0

	for i, p := range problems {
		log.Printf("Problem #%d: %s = \n", i+1, p.question)

		answerChan := make(chan string)
		go func() {
			var answer string

			fmt.Scanf("%s\n", &answer)

			answerChan <- answer
		}()

		select {
		case <-timer.C:
			log.Printf("You score is %d of %d\n", score, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.answer {
				score++
			}
		}

	}

	log.Printf("You score is %d of %d\n", score, len(problems))
	return
}

func readCsv(file *os.File) (lines [][]string) {
	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()

	eP := util.NewErrorParam(err, fmt.Sprintf("Failed to parse provide CSV.\nError: %s", err), util.Fatal)
	util.HandleError(eP)

	return lines
}

func parseLines(lines [][]string) (ret []problem) {
	ret = make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{question: line[0], answer: strings.TrimSpace(line[1])}
	}

	return ret
}
