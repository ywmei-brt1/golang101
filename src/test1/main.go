package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type InterfaceInfo struct {
	ipResult string
	nicType  string
}

func (i InterfaceInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("IpResult", i.ipResult)
	enc.AddString("NicType", i.nicType)
	return nil
}

func main() {
	// Create a logger instance
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	result := InterfaceInfo{
		ipResult: "foo",
		nicType:  "bar",
	}

	logger.Info("This is an information message",
		zap.Any("key", result),
	)
}
