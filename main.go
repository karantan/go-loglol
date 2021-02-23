package main

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logFile = "loglol.log"

// FileLogger sets up logger to a file
func FileLogger() {
	logFileLocation, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}

var logger *zap.Logger

// ZapLogger sets up a more advanced zap logger
func ZapLogger() {
	logger, _ = zap.NewProduction()
}

func zapHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

var sugarLogger *zap.SugaredLogger

// SugarLogger sets up zap logger with high-level SugaredLogger (supports formatting)
func SugarLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func sugarHTTPGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

// ZapCoreLogger creates a new logger from zapcore
func ZapCoreLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// replace with `NewJSONEncoder` to use JSON format
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	return zapcore.AddSync(file)
}

// LogInit returns a logger that shows log on the Console and a log file
func LogInit(debug bool, f *os.File) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	level := zap.InfoLevel
	if debug {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)
	logger := zap.New(core)

	return logger.Sugar()
}

// httpGetInjected is like the other HTTPGet functions except it accepts log function
func httpGetInjected(url string, log *zap.SugaredLogger) {
	log.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		log.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	// FileLogger()
	// simpleHTTPGet("www.google.com")
	// simpleHTTPGet("http://www.google.com")

	// ZapLogger()
	// defer logger.Sync()
	// zapHTTPGet("www.google.com")
	// zapHTTPGet("http://www.google.com")

	// SugarLogger()
	// defer sugarLogger.Sync()
	// sugarHTTPGet("www.google.com")
	// sugarHTTPGet("http://www.google.com")

	// ZapCoreLogger()
	// defer sugarLogger.Sync()
	// sugarHTTPGet("www.google.com")
	// sugarHTTPGet("http://www.google.com")

	home, _ := os.Getwd()
	logFileLocation, _ := os.OpenFile(home+logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log := LogInit(true, logFileLocation)
	httpGetInjected("www.google.com", log)
	httpGetInjected("http://www.google.com", log)
}
