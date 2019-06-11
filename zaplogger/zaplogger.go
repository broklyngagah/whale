package zaplogger

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger
var zapStdout *zap.Logger
var wsLogger *WriteSyncLogger

func Debug(msg string, fields ...zapcore.Field) {
	zapLogger.Debug(msg, fields...)
	if nil != zapStdout {
		zapStdout.Debug(msg, fields...)
	}
}

func Info(msg string, fields ...zapcore.Field) {
	zapLogger.Info(msg, fields...)
	if nil != zapStdout {
		zapStdout.Info(msg, fields...)
	}
}

func Warn(msg string, fields ...zapcore.Field) {
	zapLogger.Warn(msg, fields...)
	if nil != zapStdout {
		zapStdout.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zapcore.Field) {
	zapLogger.Error(msg, fields...)
	if nil != zapStdout {
		zapStdout.Error(msg, fields...)
	}
}

func Panic(msg string, fields ...zapcore.Field) {
	zapLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	zapLogger.Fatal(msg, fields...)
}

func Rotate(stdout bool) {
	wsLogger.Rotate()
}

func SetLogger(logDir, logName, logLvl string, stdout bool) bool {
	logDir, _ = filepath.Abs(logDir)
	zapLogPath := filepath.Join(logDir, logName)

	wsLogger = &WriteSyncLogger{
		Filename: zapLogPath,
	}
	w := zapcore.AddSync(wsLogger)
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.9999999"))
	}

	lvl := parseLoggerLevel(logLvl)
	zapLogger = zap.New(zapcore.NewCore(
		NewTextNoKeyEncoder(cfg),
		w,
		lvl,
	), zap.AddCallerSkip(1), zap.AddCaller())

	if stdout {
		zapStdout = zap.New(zapcore.NewCore(
			NewTextNoKeyEncoder(cfg),
			os.Stdout,
			lvl,
		), zap.AddCallerSkip(1), zap.AddCaller())
	}

	return true
}

func parseLoggerLevel(lvl string) zapcore.Level {
	var level zapcore.Level

	switch lvl {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	return level
}

func Close() {
	zapLogger.Info("Close zaplogger...")

	wsLogger.mu.Lock()
	defer wsLogger.mu.Unlock()
	wsLogger.file.Sync()
	wsLogger.close()
}
