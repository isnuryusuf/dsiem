package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func setupLogger() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	// formatter := &logrus.JSONFormatter{}
	logger.Formatter = formatter
	logger.Out = os.Stdout

	// use logrus for standard log output, those chatty 3rd-party libs ..
	log.SetOutput(logger.Writer())
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func logAndExit(err error) {
	// time.Sleep(5 * time.Minute)
	logger.Fatalf("%+v", errors.Wrap(err, ""))
}
