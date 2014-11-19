package main

import (
	"flag"
	"fmt"
	"github.com/magicshui/goutils/files"
	"io"
	"log"
	"os"
)

type debugLogging bool
type infoLogging bool
type warningLogging bool
type errorLogging bool

var (
	debug   debugLogging
	info    infoLogging
	warning warningLogging
	errl    errorLogging

	logFile     io.Writer
	logFilePath string

	debugLog   = log.New(os.Stdout, "", log.LstdFlags)
	infoLog    = debugLog
	warningLog = debugLog
	errorLog   = debugLog
)

func init() {
	flag.BoolVar((*bool)(&debug), "debug", false, "enable debug log")
	flag.BoolVar((*bool)(&info), "info", true, "enable info log")
	flag.BoolVar((*bool)(&warning), "warning", true, "enable warning log")
	flag.BoolVar((*bool)(&errl), "error", true, "enable error log")
	flag.StringVar((*string)(&logFilePath), "logfile", "", "logging to file")
}

func initLog() {
	logFile = os.Stdout
	if logFilePath != "" {
		if f, err := os.OpenFile(files.AbsPath(logFilePath),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600); err != nil {
			fmt.Printf("Can't open log file, logging to stdout: %v\n", err)
		} else {
			logFile = f
		}
	}
	log.SetOutput(logFile)

	debugLog = log.New(logFile, "[DEBUG]   ", log.LstdFlags)
	infoLog = log.New(logFile, "[INFO]    ", log.LstdFlags)
	warningLog = log.New(logFile, "[WARNING] ", log.LstdFlags)
	errorLog = log.New(logFile, "[ERROR]   ", log.LstdFlags)
}

func (d debugLogging) Printf(format string, args ...interface{}) {
	if d {
		debugLog.Printf(format, args...)
	}
}

func (d debugLogging) Println(args ...interface{}) {
	if d {
		debugLog.Println(args...)
	}
}

func (d infoLogging) Printf(format string, args ...interface{}) {
	if d {
		infoLog.Printf(format, args...)
	}
}

func (d infoLogging) Println(args ...interface{}) {
	if d {
		infoLog.Println(args...)
	}
}

func (d warningLogging) Printf(format string, args ...interface{}) {
	if d {
		warningLog.Printf(format, args...)
	}
}

func (d warningLogging) Println(args ...interface{}) {
	if d {
		warningLog.Println(args...)
	}
}

func (d errorLogging) Printf(format string, args ...interface{}) {
	if d {
		errorLog.Printf(format, args...)
	}
}

func (d errorLogging) Println(args ...interface{}) {
	if d {
		errorLog.Println(args...)
	}
}
