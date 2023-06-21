package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

func getLogLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	case "panic":
		return zapcore.PanicLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("unknown log level") // return by default
	}
}

func NewLogger(logLevel string, logPath string) (*Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	// Set log level
	level, err := getLogLevel(logLevel)
	if err != nil {
		return nil, err
	}
	cfg.Level.SetLevel(level)

	// Create log file
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	fileWriter := zapcore.AddSync(file)
	cfg.OutputPaths = []string{
		fileWriter.(*os.File).Name(),
	}

	// set log format
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.LineEnding = "\n" // Use "\n" for multi-line messages

	// set time format
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// set

	cfg.EncoderConfig = encoderConfig
	cfg.Encoding = "console"

	// Create logger
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		logger: zapLogger.WithOptions(zap.AddCallerSkip(1)), // Skip 1 level to get the caller of the logger
	}

	return logger, nil
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}
