package main

import (
	"flag"

	"github.com/google/logger"
)

const logPath = "/var/log/go/log.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	logger.Info("I'm about to do something!")
}

// func init() {
// 	flag.Parse()

// 	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
// 	if err != nil {
// 		logger.Fatalf("Failed to open log file: %v", err)
// 	}
// 	defer lf.Close()
// 	defer logger.Init("LoggerExample", *verbose, true, lf).Close()
// }
