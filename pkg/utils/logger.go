package utils

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "[DataSmith] ", log.LstdFlags|log.Lshortfile)
}
