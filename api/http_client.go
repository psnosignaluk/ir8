package api

import (
	"io"
	"net/http"
	"time"

	ir8Client "github.com/psnosignaluk/ir8/api/ir8_client"
)

type HTTPClientOptions struct {
	AppVersion        string
	CacheTTL          time.Duration
	Log               io.Writer
	LogColorize       bool
	SkipAcceptHeaders bool
	Timeout           time.Duration
}

func NewHTTPClient(opts HTTPClientOptions) (*http.Client, error) {
	clientOpts := ir8Client.ClientOptions{
		Host: "none",
	}
	return nil, nil
}
