package helper

import (
	"os"
	"testing"

	common "github.com/ci-on-dev/ci.on-cli/internal/common"
)

func TestNewHelperService(t *testing.T) {

	cacheFile := "test_helper_cache.log"
	defer os.Remove(cacheFile)

	tests := []struct {
		name        string
		level       common.LogLevel
		wantErr     bool
		wantLokiNil bool
		wantErpNil  bool
	}{
		{
			name:        "all clients provided",
			level:       common.Info,
			wantLokiNil: false,
			wantErpNil:  false,
		},
		{
			name:        "no clients provided",
			level:       common.Warn,
			wantLokiNil: true,
			wantErpNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewHelperService(tt.level)
			impl, ok := svc.(*helperService)
			if !ok {
				t.Fatalf("expected *helperService, got %T", svc)
			}
			if impl.Level != tt.level {
				t.Errorf("Level = %v; want %v", impl.Level, tt.level)
			}
		})
	}
}
