package rest

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger := zap.NewExample()
	defer logger.Sync()

	sugar := logger.Sugar()
	return sugar
}
