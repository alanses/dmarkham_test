package errors_handlers

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"test_dmarkham/blueprint/application/logger"
)

type WrappedError struct {
	Message string
	Err     error
	trace   string
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Message, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Message: info,
		Err:     err,
	}
}

func HandleError(appLogger *zap.Logger, err error) {
	if err != nil {
		appLogger.Error("something went wrong", zap.Error(err))
	}
}

func ForceHandleError(err error) {
	appLogger := logger.InitLogger(zapcore.DebugLevel)

	appLogger.Logger.Error("something went wrong", zap.Error(err))
}
