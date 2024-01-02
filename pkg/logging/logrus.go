package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0744)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	switch os.Getenv("LOG_LEVEL") {
	case Trace:
		l.SetLevel(logrus.TraceLevel)
	case Debug:
		l.SetLevel(logrus.DebugLevel)
	case Info:
		l.SetLevel(logrus.InfoLevel)
	case Warn:
		l.SetLevel(logrus.WarnLevel)
	case Error:
		l.SetLevel(logrus.ErrorLevel)
	default:
		l.SetLevel(logrus.DebugLevel)
	}

	e = logrus.NewEntry(l)
}

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (h *writerHook) Levels() []logrus.Level {
	return h.LogLevels
}

func (h *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, writer := range h.Writer {
		_, err = writer.Write([]byte(line))
	}
	return err
}
