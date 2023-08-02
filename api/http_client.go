package api

import (
	"io"
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

func ShortGet(url string, client *http.Client) (map[string]string, error) {
	m := make(map[string]string)

	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	_, err = io.ReadAll(resp.Body)
	elapsed := time.Since(start)

	defer resp.Body.Close()

	m["status"] = resp.Status
	m["elapsed"] = elapsed.String()

	if err != nil {
		return m, err
	}

	return m, err
}
