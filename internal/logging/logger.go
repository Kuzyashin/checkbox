package logging

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Caller().Timestamp()
	return logger.Logger()
}
