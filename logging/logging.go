package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new sugar logger that writes to the console
func New(name string, debug bool) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	level := zap.InfoLevel
	if debug {
		level = zap.DebugLevel
	}

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
	logger := zap.New(core, zap.AddCaller())

	defer logger.Sync() // flush the log

	return logger.Named(name).Sugar()
}

// NewFileLogger creates a new sugar logger that writes to a file
func NewFileLogger(path string, name string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{path}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync() // flush the log

	return logger.Named(name).Sugar(), nil
}
