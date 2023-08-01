package api

import (
	"io"
)

type IR8ClientOptions struct {
	AppVersion  string
	Host        string
	Log         io.Writer
	LogColorize bool
	UserAgent   string
}
