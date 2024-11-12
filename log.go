package tools

import (
	"github.com/heybrovn/tools/log"
	"os"
)

var Logger log.Logger

func init() {
	logLevel := log.Info // os.Getenv("LOGLEVEL")
	if os.Getenv("LOGLEVEL") == "debug" {
		logLevel = log.Debug
	}

	//setup logger
	logOpts := log.Options{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      logLevel, // "debug",
		EnableFile:        false,
	}
	Logger, _ = log.New(log.ZapLogger, logOpts)
}
