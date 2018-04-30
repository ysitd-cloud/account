package providers

import (
	"os"

	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
	"golang.ysitd.cloud/log"
)

var Logger *logrus.Logger

func initLogger() *logrus.Logger {
	if Logger == nil {
		if logFile := os.Getenv("LOG_FILE"); logFile != "" {
			var err error
			Logger, err = log.NewForContainer(logFile)
			if err != nil {
				panic(err)
			}
		} else {
			Logger = logrus.New()
		}

		if os.Getenv("VERBOSE") != "" {
			Logger.SetLevel(logrus.DebugLevel)
		}
	}
	return Logger
}

func InjectLogger(graph *inject.Graph) {
	logger := initLogger()
	graph.Provide(
		NewObject(logger),
		NewNamedObject("osin logger", logger.WithField("source", "osin")),
	)
}
