package logger

import (
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	debug    bool
	shellLog *log.Logger
	fileLog  *log.Logger
}

func New(debug bool, filename string) *Logger {
	shellLogger := log.New(os.Stdout, "", 0)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		shellLogger.SetPrefix("DEBUG\t")
		shellLogger.SetFlags(log.Ldate | log.Ltime)
		shellLogger.Printf("couldn't open file %s, error %v", filename, err)
	}

	fileLogger := log.New(file, "", 0)

	return &Logger{
		debug:    debug,
		shellLog: shellLogger,
		fileLog:  fileLogger,
	}
}

func (l *Logger) Println(level LogLevel, msg string) {
	switch level {
	case DEBUG:
		if l.debug == true {
			l.shellLog.SetPrefix("DEBUG\t")
			l.shellLog.SetFlags(log.Ldate | log.Ltime)
			l.shellLog.Println(msg)
		}
		break

	case INFO:
		l.fileLog.SetPrefix("INFO\t")
		l.fileLog.SetFlags(log.Ldate | log.Ltime)
		l.fileLog.Println(msg)

	case WARNING:
		l.shellLog.SetPrefix("WARNING\t")
		l.shellLog.SetFlags(log.Ldate | log.Ltime)
		l.shellLog.Println(msg)

	case ERROR:
		l.fileLog.SetPrefix("ERROR\t")
		l.fileLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		l.fileLog.Println(msg)

	case FATAL:
		l.fileLog.SetPrefix("FATAL\t")
		l.fileLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		l.fileLog.Println(msg)
		os.Exit(1)
	}

}
