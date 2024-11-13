package utils

import (
	"log"
	"os"
)

type LogErrorData struct {
	TimeStamp          int64                  `json:"timeStamp"`
	ErrorMessage       string                 `json:"errorMessage"`
	ErrorCodeForClient string                 `json:"errorCodeForClient"`
	ErrorOrigin        string                 `json:"errorOrigin"`
	ErrorLevel         int                    `json:"errorLevel"`
	ErrorStack         error                  `json:"errorStack"`
	InputParams        map[string]interface{} `json:"inputParams"`
}

func LogError(errorMessage string, errorCodeForClient string, errorLevel int, errorStack error, inputParams map[string]interface{}) {
	errorOrigin := os.Getenv("SERVICE_NAME")
	LogErrorInfo := LogErrorData{
		TimeStamp:          GetCurrentTime().Unix(),
		ErrorMessage:       errorMessage,
		ErrorCodeForClient: errorCodeForClient,
		ErrorOrigin:        errorOrigin,
		ErrorLevel:         errorLevel,
		ErrorStack:         errorStack,
	}
	if inputParams != nil {
		LogErrorInfo.InputParams = inputParams
	}
	log.Println(LogErrorInfo)
	if errorMessage == "gocql: no hosts available in the pool" {
		os.Exit(1)
	}
}
