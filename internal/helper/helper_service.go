package helper

import (
	"fmt"
	"sync"
	"time"

	common "github.com/ci-on-dev/ci.on-cli/internal/common"
	"github.com/ci-on-dev/ci.on-cli/internal/utils"
)

// PrintWithColor allows colored console output; can be overridden in tests.
var PrintWithColor = utils.PrintWithColor

// HelperService defines the public API.
type HelperService interface {
	Print(message string, level common.LogLevel)
}

// helperService is the concrete implementation with in-memory and on-disk cache.
type helperService struct {
	Level         common.LogLevel
	buffer        [][]byte    // in-memory buffer of failed messages
	mu            *sync.Mutex // protects buffer
	stopChan      chan struct{}
	retryInterval time.Duration
	done          *sync.WaitGroup
}

// NewHelperService constructs the service, starts the retry loop, and returns it.
// cacheFilePath is the path to the disk cache file for pending logs.
func NewHelperService(
	Level common.LogLevel,
) HelperService {
	mu := &sync.Mutex{}
	done := &sync.WaitGroup{}
	retryInterval := 10 * time.Second

	svc := &helperService{
		Level:         Level,
		buffer:        make([][]byte, 0),
		mu:            mu,
		stopChan:      make(chan struct{}),
		retryInterval: retryInterval,
		done:          done,
	}

	done.Add(1)
	return svc
}

// Print attempts to send the message immediately to remote clients.
// On failure, it buffers in memory (and will be retried).
// Always prints to console if level passes threshold.
func (svc *helperService) Print(message string, level common.LogLevel) {
	fmt.Printf("svc.Level: %s\n", svc.Level)
	fmt.Printf("level: %s\n", level)
	if svc.Level < level {
		return
	}

	// always print locally
	PrintWithColor(level, message)
}
