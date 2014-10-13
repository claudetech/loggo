package loggo_default

import (
	"github.com/claudetech/loggo"
	"os"
	"strings"
)

// Defaults logger to use for simple cases
var Log *loggo.Logger = func() *loggo.Logger {
	name := strings.TrimLeft(strings.ToUpper(os.Args[0]), "./")
	log := loggo.New(name)
	log.AddAppenderWithFilter(loggo.NewStdoutAppender(), &loggo.MaxLogLevelFilter{MaxLevel: loggo.Info}, loggo.Color)
	log.AddAppenderWithFilter(loggo.NewStderrAppender(), &loggo.MinLogLevelFilter{MinLevel: loggo.Warning}, loggo.Color)
	return log
}()
