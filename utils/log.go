package utils

import "go.uber.org/zap"

var Logger *zap.Logger
var err error

func init() {
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
