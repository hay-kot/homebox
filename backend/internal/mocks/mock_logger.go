package mocks

import (
	"os"

	"github.com/hay-kot/content/backend/pkgs/logger"
)

func GetStructLogger() *logger.Logger {
	return logger.New(os.Stdout, logger.LevelDebug)
}
