package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/hhhhp52/webtest/src/utils/config"
)

var (
	logger             *log.Logger
	file               *os.File
	loggerRegexp       = regexp.MustCompile(`src/utils/logger/.*.go`)
	middlewareRegexp   = regexp.MustCompile(`src/middleware/.*.go`)
	errorHanlderRegexp = regexp.MustCompile(`src/handler/.*Handler.go`)
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	logPath := fmt.Sprintf("%s/../../../%s%s", dir, config.Get("log.path"), config.Get("log.file"))
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.SetPrefix("[ ")
		log.Fatalln(fmt.Sprintf("] %s - %s : %s", "ERROR", callStack(), err))
	} else {
		logger = log.New(logFile, "", log.LstdFlags)
		logger.SetPrefix("[ ")
	}
}

// Close should be called before main process ends.
func Close() {
	if file != nil {
		file.Close()
	}
	file, logger = nil, nil
}

// Debug outputs the message into log file
func Debug(msg interface{}) {
	output("DEBUG", msg)
}

// Info outputs the message into log file
func Info(msg interface{}) {
	output("INFO", msg)
}

// Warn outputs the message into log file
func Warn(msg interface{}) {
	output("WARN", msg)
}

// Error outputs the message into log file
func Error(msg interface{}) {
	output("ERROR", msg)
}

func output(level string, msg interface{}) {
	logger.Println(fmt.Sprintf("] %s - %s : %v", level, callStack(), msg))
}

func callStack() string {
	resul := ""
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)

		if !ok {
			break
		} else if !loggerRegexp.MatchString(file) && !middlewareRegexp.MatchString(file) && !errorHanlderRegexp.MatchString(file) && strings.Contains(file, "auroratechit-aurora-payment/") {
			splittedPath := strings.Split(file, "auroratechit-aurora-payment/")
			relativePath := splittedPath[len(splittedPath)-1]
			resul = fmt.Sprintf("File \"%v\", line %v", relativePath, line)
		}
	}
	return resul
}

func fullStack() []string {
	var stack []string
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)

		if !ok {
			break
		} else if !loggerRegexp.MatchString(file) && !middlewareRegexp.MatchString(file) && !errorHanlderRegexp.MatchString(file) && strings.Contains(file, "auroratechit-aurora-payment/") {
			splittedPath := strings.Split(file, "auroratechit-aurora-payment/")
			relativePath := splittedPath[len(splittedPath)-1]
			stack = append(stack, fmt.Sprintf("File \"%v\", line %v", relativePath, line))
		}
	}
	return stack
}
