package logger

import (
	"golang-agnostic-template/src/pkg/config"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ILogger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	With(fields ...zap.Field) *zap.Logger
	Sync() error
}

type zapLogger struct {
	logger *zap.Logger
}

var instance ILogger
var once sync.Once

func New() (ILogger, error) {
	var initErr error
	once.Do(func() {
		instance, initErr = newZapLogger()
	})
	return instance, initErr
}

func newZapLogger() (ILogger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Params.FilePath,
		MaxSize:    config.Params.FileSize,     // megabytes
		MaxBackups: config.Params.MaxBackups,   // number of log files
		MaxAge:     config.Params.MaxDuration,  // days
		Compress:   config.Params.FileCompress, // disabled by default
	}

	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(config.Params.Level)); err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberjackLogger),
		zapcore.AddSync(os.Stdout),
	)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // o NewConsoleEncoder para dev
		writeSyncer,
		level,
	)

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}

	if config.Params.Environment == "dev" {
		options = append(options, zap.Development())
	}

	return &zapLogger{
		logger: zap.New(core, options...),
	}, nil
}

// --- Implementación de la interfaz ---
func (l *zapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}

func (l *zapLogger) With(fields ...zap.Field) *zap.Logger {
	return l.logger.With(fields...)
}
