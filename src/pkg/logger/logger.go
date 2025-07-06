package logger

import (
	"os"
	"sync"

	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerField struct {
	Key   string
	Value interface{}
}

type ILogger interface {
	Debug(msg string, fields ...LoggerField)
	Info(msg string, fields ...LoggerField)
	Warn(msg string, fields ...LoggerField)
	Error(msg string, fields ...LoggerField)
	Fatal(msg string, fields ...LoggerField)
	With(fields ...LoggerField) ILogger
	Sync() error
}

type zapLogger struct {
	logger *zap.Logger
}

var (
	instance ILogger
	once     sync.Once
)

func NewLogger() (ILogger, error) {
	var initErr error
	once.Do(func() {
		instance, initErr = newZapLogger()
	})
	if initErr != nil {
		return nil, initErr
	}
	return instance, nil
}

func newZapLogger() (ILogger, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Params.FilePath,
		MaxSize:    config.Params.FileSize,
		MaxBackups: config.Params.MaxBackups,
		MaxAge:     config.Params.MaxDuration,
		Compress:   config.Params.FileCompress,
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
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var writeSyncer zapcore.WriteSyncer
	if config.Params.FilePath != "" {
		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(lumberjackLogger),
			zapcore.AddSync(os.Stdout),
		)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		level,
	)

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}

	if config.Params.Environment == utils.DEV {
		options = append(options, zap.Development())
	}

	logger := zap.New(core, options...)
	return &zapLogger{logger: logger}, nil
}

func toZapFields(fields ...LoggerField) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		switch v := f.Value.(type) {
		case error:
			zapFields[i] = zap.Error(v)
		default:
			zapFields[i] = zap.Any(f.Key, f.Value)
		}
	}
	return zapFields
}

func (l *zapLogger) Debug(msg string, fields ...LoggerField) {
	l.logger.Debug(msg, toZapFields(fields...)...)
}

func (l *zapLogger) Info(msg string, fields ...LoggerField) {
	l.logger.Info(msg, toZapFields(fields...)...)
}

func (l *zapLogger) Warn(msg string, fields ...LoggerField) {
	l.logger.Warn(msg, toZapFields(fields...)...)
}

func (l *zapLogger) Error(msg string, fields ...LoggerField) {
	l.logger.Error(msg, toZapFields(fields...)...)
}

func (l *zapLogger) Fatal(msg string, fields ...LoggerField) {
	l.logger.Fatal(msg, toZapFields(fields...)...)
}

func (l *zapLogger) With(fields ...LoggerField) ILogger {
	return &zapLogger{logger: l.logger.With(toZapFields(fields...)...)}
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}
