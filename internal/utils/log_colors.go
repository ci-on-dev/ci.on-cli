package utils

import common "github.com/ci-on-dev/ci.on-cli/internal/common"

var Reset = "\033[0m"
var Error = "\033[31m"
var Warn = "\033[33m"
var Info = "\033[34m"
var Debug = "\033[35m"

func GetLogColor(level common.LogLevel) string {
	switch level {
	case common.Error:
		return Error
	case common.Warn:
		return Warn
	case common.Info:
		return Info
	case common.Debug:
		return Debug
	default:
		return Reset
	}
}
