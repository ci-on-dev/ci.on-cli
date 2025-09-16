package logs

import (
	"fmt"
	"os"
	"strings"

	common "github.com/ci-on-dev/ci.on-cli/internal/common"
	"github.com/ci-on-dev/ci.on-cli/internal/helper"
)

type LogsService interface {
	Debugf(format string, v ...any)
	Debug(message string)
	Infof(format string, v ...any)
	Info(message string)
	Warnf(format string, v ...any)
	Warn(message string)
	Errorf(statusCode int, errorCode string, details string, v ...any)
	Error(statusCode int, errorCode string, details string)
	Fatalf(statusCode int, errorCode string, details string, v ...any)
	Fatal(statusCode int, errorCode string, details string)
}

type logsService struct {
	Level         common.LogLevel
	helperService helper.HelperService
}

func NewLogsService(level string) LogsService {
	var helperService helper.HelperService

	cacheFile := "common_helper_cache.log"
	defer os.Remove(cacheFile)

	_logLevel := common.Info
	switch strings.ToLower(level) {
	case "fatal":
		_logLevel = common.Fatal
	case "warn":
		_logLevel = common.Warn
	case "debug":
		_logLevel = common.Debug
	case "info":
		_logLevel = common.Info
	case "error":
		_logLevel = common.Error
	default:
		_logLevel = common.Info
	}
	fmt.Printf("logLevel 2: %s\n", _logLevel)

	helperService = helper.NewHelperService(_logLevel)
	return logsService{Level: _logLevel, helperService: helperService}
}
