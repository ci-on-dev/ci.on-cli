package utils

import (
	"fmt"
	"log"

	"strings"

	common "github.com/ci-on-dev/ci.on-cli/internal/common"
)

func PrintWithColor(level common.LogLevel, message string) {
	logSuffix := strings.ToUpper(fmt.Sprintf("[%s] ", level.String()))
	switch level {
	case common.Error:
		log.Println(Error + logSuffix + message + Reset)
	case common.Warn:
		log.Println(Warn + logSuffix + message + Reset)
	case common.Info:
		log.Println(Info + logSuffix + message + Reset)
	case common.Debug:
		log.Println(Debug + logSuffix + message + Reset)
	default:
		log.Println(Reset + message)
	}
}
