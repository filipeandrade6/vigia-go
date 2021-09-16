// Package logger provides a convience function to contruction a logger for use.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New contructs a Sugared Logger that writes to stdout and
// provides human readable timestamps.
func New(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
