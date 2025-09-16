package logs

type LogLevel int

const (
	Debug LogLevel = 4
	Info  LogLevel = 3
	Warn  LogLevel = 2
	Error LogLevel = 1
	Fatal LogLevel = 0
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return "unknown"
	}
}
