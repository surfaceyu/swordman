package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"gitee.com/surfaceyu/logrusformator"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	level   = "DEBUG"
	pattern = "%Y-%m-%d-%H"
	maxAge  = 30
	// maxAgeComment = "days"
	rotate = 12
	// rotateComment = "hours"
	filename = "./logs/word-game"
	log      = logrus.New()
)

func init() {
	logger, _ := rotatelogs.New(
		fmt.Sprintf("%s.%s", filename, pattern),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(rotate)*time.Hour),
	)

	mw := io.MultiWriter(os.Stdout, logger)
	log.SetFormatter(logrusformator.PomeloFormator)
	log.SetOutput(mw)
	log.SetReportCaller(true)
	setLevel(level)
}

func setLevel(level string) {
	switch strings.ToUpper(level) {
	default:
		log.SetLevel(logrus.InfoLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	}
}

func GetLogger() *logrus.Logger {
	return log
}

var Debug = log.Debug
var Debugln = log.Debugln
var Debugf = log.Debugf
var Info = log.Info
var Infoln = log.Infoln
var Infof = log.Infof
var Warn = log.Warn
var Warnln = log.Warnln
var Warnf = log.Warnf
var Error = log.Error
var Errorln = log.Errorln
var Errorf = log.Errorf
var Fatal = log.Fatal
var Fatalln = log.Fatalln
var Fatalf = log.Fatalf
