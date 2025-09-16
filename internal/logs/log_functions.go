package logs

import (
	"fmt"
	"os"

	common "github.com/ci-on-dev/ci.on-cli/internal/common"
)

func (service logsService) Debugf(format string, v ...any) {
	level := common.Debug
	message := fmt.Sprintf(format, v...)
	service.helperService.Print(message, level)
}

func (service logsService) Debug(message string) {
	level := common.Debug
	service.helperService.Print(message, level)
}

func (service logsService) Infof(format string, v ...any) {
	level := common.Info
	message := fmt.Sprintf(format, v...)
	service.helperService.Print(message, level)
}

func (service logsService) Info(message string) {
	level := common.Info
	service.helperService.Print(message, level)
}

func (service logsService) Warn(message string) {
	level := common.Warn
	service.helperService.Print(message, level)
}

func (service logsService) Warnf(format string, v ...any) {
	level := common.Warn
	message := fmt.Sprintf(format, v...)
	service.helperService.Print(message, level)
}

func (service logsService) Error(statusCode int, errorCode string, details string) {
	level := common.Error
	message := fmt.Sprintf("%d %s %s", statusCode, errorCode, details)

	service.helperService.Print(message, level)
}

func (service logsService) Errorf(statusCode int, errorCode string, details string, v ...any) {
	level := common.Error
	message := fmt.Sprintf("%d %s %s", statusCode, errorCode, details)

	service.helperService.Print(message, level)
}

func (service logsService) Fatalf(statusCode int, errorCode string, details string, v ...any) {
	level := common.Error
	message := fmt.Sprintf("%d %s %s", statusCode, errorCode, details)

	service.helperService.Print(message, level)
	os.Exit(1)
}

func (service logsService) Fatal(statusCode int, errorCode string, details string) {
	level := common.Error
	message := fmt.Sprintf("%d %s %s", statusCode, errorCode, details)

	service.helperService.Print(message, level)
	os.Exit(1)
}
