package api

import (
	"io"
	"time"
)

type DNSClientOptions struct {
	AppVersion  string
	Log         io.Writer
	LogColorize bool
	Timeout     time.Duration
}
