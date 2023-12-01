package main

import (
	"go.uber.org/zap"
)

type Foo struct {
	Name string
	Age  int
}

func main() {
	// Create a logger instance
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	f := Foo{"foo", 100}

	logger.Info("This is an information message",
		zap.Any("key", f),
	)
}
