package logging

import "go.uber.org/zap"

var Logger *zap.Logger

func ConfigureLogger() error {
	var err error
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	Logger, err = config.Build()
	return err
}
