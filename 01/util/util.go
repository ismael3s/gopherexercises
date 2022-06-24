package util

import "log"

type ErrorFlow string

const (
	LogOnly ErrorFlow = "LOG-ONLY"
	Fatal   ErrorFlow = "FATAL"
)

type ErrorParam struct {
	err     error
	message string
	flow    ErrorFlow
}

func NewErrorParam(err error, message string, flow ErrorFlow) *ErrorParam {
	return &ErrorParam{err: err, message: message, flow: flow}
}

func HandleError(params *ErrorParam) {
	if params.err == nil {
		return
	}

	switch params.flow {
	case Fatal:
		log.Fatalln(params.message)
	default:
		log.Println(params.message)
	}
}
