package constants

type LogMode string

var (
	logMode LogMode = LogModeDefault
)

const (
	LogModeDefault = "default"
	LogModeVerbose = "verbose"
)

func SetLogMode(mode string) {
	logMode = LogMode(mode)
}

// GetLogMode returns the current log mode
func GetLogMode() LogMode {
	return logMode
}
