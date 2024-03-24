package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	var err error
	config := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.StacktraceKey = ""
	zapcore.TimeEncoderOfLayout("Jan _2 15:04:05.000000000")
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	defer logger.Sync() // flushes buffer, if any
	Logger = logger.Sugar()
	if err != nil {
		panic(err)
	}
}
