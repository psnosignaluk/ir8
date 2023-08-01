package api

import (
	"io"
	"time"
)

type HTTPClientOptions struct {
	AppVersion        string
	CacheTTL          time.Duration
	Log               io.Writer
	LogColorize       bool
	SkipAcceptHeaders bool
	Timeout           time.Duration
}
