package api

import (
	"net"
	"net/http"
	"time"
)

type HTTPClientOptions struct {
	Host        string
	KeepAlive   time.Duration
	DialTimeout time.Duration
	Timeout     time.Duration
	TLSTimeout  time.Duration
}

func NewHTTPClient(opts HTTPClientOptions) (*http.Client, error) {
	transport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   opts.DialTimeout,
			KeepAlive: opts.KeepAlive,
		}).Dial,
		TLSHandshakeTimeout: opts.TLSTimeout,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   opts.Timeout,
	}

	return &client, nil
}
