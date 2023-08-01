package api

import (
	"io"
	"net/http"
	"time"
)

type HTTPClientOptions struct {
	AppVersion        string
	Host              string
	CacheTTL          time.Duration
	Log               io.Writer
	LogColorize       bool
	SkipAcceptHeaders bool
	Timeout           time.Duration
}

func NewHTTPClient(opts HTTPClientOptions) (*http.Client, error) {
	return nil, nil
}
