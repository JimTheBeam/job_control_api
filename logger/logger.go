package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger initialize logrus logger
func InitLogger(log *logrus.Logger, lvl string) {
	log.SetFormatter(&logrus.TextFormatter{
		// ForceColors: true,
		DisableColors:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	log.SetOutput(os.Stdout)

	setLogLevel(log, lvl)

}

// SetLogLevel - set log level, default warning level
func setLogLevel(log *logrus.Logger, lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		log.Errorf("can't parse log level: %v", err)
		// default log level
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	log.Infof("set log level: %s", level.String())
}

// SetBaseFields - set base fields to log
func SetBaseFields(log *logrus.Logger, p, f string) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"package":  p,
		"function": f,
	})
}
