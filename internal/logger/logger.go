package logger

import (
	"log"
	"os"
)

type CustomLoggerStruct struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

var CustomLogger = CustomLoggerStruct{
	Info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	Warn:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime),
	Error: log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime),
}
