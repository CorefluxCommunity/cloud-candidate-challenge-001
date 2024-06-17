package loggers

import (
	"log"
	"os"
)

var (
    LogError = log.New(os.Stderr, "", log.LstdFlags)
    LogEvent = log.New(os.Stdout, "", log.LstdFlags)
)