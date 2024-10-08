package internal

import (
    "go.uber.org/zap"
    "log"
)

var Logger *zap.Logger

func InitLogger() {
    var err error

    Logger, err = zap.NewProduction()
    if err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }
}
