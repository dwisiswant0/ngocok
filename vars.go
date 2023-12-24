package main

import (
	"log/slog"

	"github.com/henvic/httpretty"
)

var (
	ID string

	httplog *httpretty.Logger
	log     *slog.Logger
	opt     *options

	skipHeaders = []string{
		"X-Forwarded-For",
		"X-Forwarded-Host",
		"X-Forwarded-Proto",
	}
)

var responseHeaders = map[string]string{
	"Content-Type": "text/plain; charset=utf-8",
	"Server":       AppName + "/" + AppVersion,

	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "*",
	"Access-Control-Allow-Origin":      "*",
}
