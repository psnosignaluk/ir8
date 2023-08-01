package api

import (
	"io"
	"time"
)

type K8SClientOptions struct {
	AppVersion  string
	Log         io.Writer
	LogColorize bool
	Timeout     time.Duration
}
