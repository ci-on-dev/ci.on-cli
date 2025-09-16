package logs

import (
	"testing"
)

func TestLogLevelString(t *testing.T) {
	testCases := []struct {
		level    LogLevel
		expected string
	}{
		{Debug, "debug"},
		{Info, "info"},
		{Warn, "warn"},
		{Error, "error"},
		{LogLevel(5), "unknown"},
		{LogLevel(0), "unknown"},
		{LogLevel(-1), "unknown"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			if result := tc.level.String(); result != tc.expected {
				t.Errorf("Expected %s for LogLevel %d, got %s", tc.expected, tc.level, result)
			}
		})
	}
}

func TestLogLevelConstants(t *testing.T) {
	if Debug != 4 {
		t.Errorf("Debug should be 4, got %d", Debug)
	}
	if Info != 3 {
		t.Errorf("Info should be 3, got %d", Info)
	}
	if Warn != 2 {
		t.Errorf("Warn should be 2, got %d", Warn)
	}
	if Error != 1 {
		t.Errorf("Error should be 1, got %d", Error)
	}
}
