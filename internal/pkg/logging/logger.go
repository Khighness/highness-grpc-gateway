package logging

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// @Author Chen Zikang
// @Email  zikang.chen@shopee.com
// @Since  2022-09-08

// InitLogger initializes zap.Logger then we can use  zap.L() to get this logger.
// If level is Debug, then log will be printed to console and also wrote to file.
// Otherwise, log will be only wrote to file.
func InitLogger(level zapcore.Level) {
	var core zapcore.Core
	fileCore := zapcore.NewCore(zapFileEncoder(), zapWriteSyncer(), zapLevelEnabler())
	if level == zapcore.DebugLevel {
		consoleCore := zapcore.NewCore(zapConsoleEncoder(), os.Stdout, level)
		core = zapcore.NewTee(fileCore, consoleCore)
	} else {
		core = zapcore.NewTee(fileCore)
	}
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func zapLevelEnabler() zapcore.Level {
	return zapcore.InfoLevel
}

func zapEncodeConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		CallerKey:        "caller_line",
		FunctionKey:      zapcore.OmitKey,
		StacktraceKey:    "stacktrace",
		LineEnding:       "\n",
		EncodeLevel:      zapEncodeLevel,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.MillisDurationEncoder,
		EncodeCaller:     zapEncodeCaller,
		ConsoleSeparator: " ",
	}
}

func zapFileEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapEncodeConfig())
}

func zapConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapEncodeConfig())
}

func zapEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func zapEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}

func zapWriteSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxSize:    1000,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
