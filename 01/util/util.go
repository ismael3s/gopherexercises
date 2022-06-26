package util

import "log"

type ErrorFlow string

const (
	LogOnly ErrorFlow = "LOG-ONLY"
	Fatal   ErrorFlow = "FATAL"
)

type ErrorParam struct {
	Err     error
	Message string
	Flow    ErrorFlow
}

func NewErrorParam(err error, message string, flow ErrorFlow) *ErrorParam {
	return &ErrorParam{Err: err, Message: message, Flow: flow}
}

func HandleError(params *ErrorParam) {
	if params.Err == nil {
		return
	}

	switch params.Flow {
	case Fatal:
		log.Fatalln(params.Message)
	default:
		log.Println(params.Message)
	}
}
