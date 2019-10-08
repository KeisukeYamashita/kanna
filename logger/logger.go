package logger

import "go.uber.org/zap"

// Logger ...
type Logger struct {
	Logger  *zap.Logger
	Verbose bool
}

// Info ...
func (logger *Logger) Info(str string, fields ...zap.Field) {
	if logger.Verbose {
		logger.Logger.Info(str, fields...)
	}
}

// Fatal ...
func (logger *Logger) Fatal(str string, fields ...zap.Field) {
	if logger.Verbose {
		logger.Logger.Fatal(str, fields...)
	}
}
