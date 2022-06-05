package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createDirectoryIfNotExists() {
	path, _ := os.Getwd()

	_, err := os.Stat(fmt.Sprintf("%s/logs", path))
	if err != nil {
		log.Panic(err)
		return
	}
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", os.ModePerm)
	}
}

func getLogWriter() zapcore.WriteSyncer {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path+"/logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	createDirectoryIfNotExists()

	writerSync := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSync, zapcore.DebugLevel)
	logg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logg)
}
