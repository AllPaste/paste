package log

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

func (l *Logger) Debug(args ...any) {
	l.loger.Sugar().Debug(args...)
}

func (l *Logger) Info(args ...any) {
	l.loger.Sugar().Info(args...)
}

func (l *Logger) Warn(args ...any) {
	l.loger.Sugar().Warn(args...)
}

func (l *Logger) Error(args ...any) {
	l.loger.Sugar().Error(args...)
}

func (l *Logger) Fatal(args ...any) {
	l.loger.Sugar().Fatal(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.loger.Sugar().Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.loger.Sugar().Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.loger.Sugar().Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.loger.Sugar().Errorf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.loger.Sugar().Fatalf(template, args...)
}

type Logger struct {
	loger *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

var std = New(os.Stderr, InfoLevel)

func Default() *Logger {
	return std
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		CallerKey:     "caller",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
		}, // 自定义时间格式
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			zapcore.CapitalColorLevelEncoder(level, enc)
			// enc.AppendString("[" + level.CapitalString() + "]")
		}, // 小写编码器
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + caller.TrimmedPath() + "]")
		}, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(writer),
		level,
	)

	logger := &Logger{
		loger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)),
		level: level,
	}

	return logger
}

func (l *Logger) Sync() error {
	return l.loger.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

// ZapLogger return zap.Logger object.
func (l *Logger) ZapLogger() *zap.Logger {
	return l.loger
}
