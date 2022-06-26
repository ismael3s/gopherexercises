package entities

type Problem struct {
	Question string
	Answer   string
}

func NewProblem(question string, answer string) Problem {
	return Problem{Question: question, Answer: answer}
}
