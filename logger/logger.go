package logger

import (
    "go.uber.org/zap"
)

// LoggerWrapper wraps the zap logger for easier use
type LoggerWrapper struct {
    logger *zap.SugaredLogger
}

// NewLoggerWrapper initializes the zap logger
func NewLoggerWrapper() *LoggerWrapper {
    zapLogger, err := zap.NewDevelopment() // Use zap.NewProduction() for production environments
    if err != nil {
        panic("Failed to initialize zap logger: " + err.Error())
    }

    sugarLogger := zapLogger.Sugar()
    return &LoggerWrapper{logger: sugarLogger}
}

// Info logs informational messages
func (lw *LoggerWrapper) Info(message string, fields ...interface{}) {
    lw.logger.Infow(message, fields...)
}

// Debug logs debug messages
func (lw *LoggerWrapper) Debug(message string, fields ...interface{}) {
    lw.logger.Debugw(message, fields...)
}

// Error logs error messages
func (lw *LoggerWrapper) Error(message string, fields ...interface{}) {
    lw.logger.Errorw(message, fields...)
}

// Close flushes the logger to ensure all logs are written
func (lw *LoggerWrapper) Close() {
    lw.logger.Sync()
}
