package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ismael3s/go-01/entities"
	"github.com/ismael3s/go-01/factories"
)

func main() {
	timeLimit := 3

	problemRepository := factories.NewProblemRepository(factories.POSTGRES)

	problems := problemRepository.Read()

	startSolving(problems, timeLimit)
}

func startSolving(problems []entities.Problem, timeLimit int) (score int) {
	timer := time.NewTimer(time.Second * time.Duration(timeLimit))
	score = 0

	for i, p := range problems {
		log.Printf("Problem #%d: %s = \n", i+1, p.Question)

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
			if answer == p.Answer {
				score++
			}
		}

	}

	log.Printf("You score is %d of %d\n", score, len(problems))
	return
}
