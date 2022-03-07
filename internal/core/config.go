package core

import (
	"github.com/pedrocmart/canvas/internal/core/log"
)

type ContextKey string

const (
	// ServiceName const
	ServiceName string = "ASCII canvas"
	// Version const
	Version string = "1.0.0"

	ContextKeyRequestTime ContextKey = "request_time"
	ContextKeyRequestID   ContextKey = "requestID"

	CanvasSizeRows    int = 100
	CanvasSizeColumns int = 100
)

type Config struct {
	Log  *log.Logger
	Env  string
	Host string
}
