package utility

import (
	"fmt"
	"log/slog"
	"os"
)

func InitializeLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // setting the output path for logger
	slog.SetDefault(logger)
	return logger // setting global logger
}

func PrintError(caller, message string, err error) {
	slog.Error(fmt.Sprintf("%v : %v", caller, message), "Error", err.Error())
}


