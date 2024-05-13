package log

import (
	"github.com/xusenlin/timeSpliceFile"
	"log/slog"
)

var Slog *slog.Logger
var GIN *slog.Logger

func InitLogger(dir string) error {

	logFile, err := timeSpliceFile.New(dir, timeSpliceFile.SplitIntervalDay, "log")
	if err != nil {
		return err
	}
	textHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	Slog = slog.New(textHandler)

	ginLogFile, err := timeSpliceFile.New(dir, timeSpliceFile.SplitIntervalDay, "gin.log")
	if err != nil {
		return err
	}
	ginTextHandler := slog.NewTextHandler(ginLogFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	GIN = slog.New(ginTextHandler)
	return nil
}
