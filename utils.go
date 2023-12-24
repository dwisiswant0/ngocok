package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"net/http"
	"net/url"
	"path/filepath"

	"github.com/dwisiswant0/clientip"
	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func (r *runner) logreq(req *http.Request) {
	ip := clientip.FromRequest(req)

	log.Debug(reqReceived, "addr", ip)

	if r.File != nil {
		header := fmt.Sprintf(
			"* %s from %s @ %s",
			reqReceived, ip.String(),
			time.Now().Format(time.RFC3339),
		)

		_, err := r.File.WriteString(header + "\n")
		if err != nil {
			log.Warn("cannot write string to a file", "err", err, "file", r.File.Name())
		}
	}

	httplog.PrintRequest(req)
}

func (r *runner) getTunnel() (config.Tunnel, error) {
	endpoint := strings.ToLower(r.opt.Endpoint)

	switch endpoint {
	case "http":
		return config.HTTPEndpoint(), nil
	case "tcp":
		return config.TCPEndpoint(), nil
	case "":
		return nil, errors.New(errGetTunnelEmpty)
	default:
		return nil, fmt.Errorf(errGetTunnelCustom, endpoint)
	}
}

func parseURL(uri string) (string, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}

	if u.Scheme == "tcp" {
		u.Scheme = "http"
	}

	return u.String(), nil
}

func (r *runner) getAuthtokenOpt() (ngrok.ConnectOption, error) {
	if r.opt.Authtoken != "" {
		return ngrok.WithAuthtoken(r.opt.Authtoken), nil
	}

	if os.Getenv(ngrokAuthtokenEnv) != "" {
		return ngrok.WithAuthtokenFromEnv(), nil
	}

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	cfgDir = filepath.Join(cfgDir, "ngrok")

	viper.SetConfigName("ngrok")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgDir)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.New(errGetAuthtokenOpt)
	}

	return ngrok.WithAuthtoken(viper.GetString(ngrokAuthtoken)), nil
}
