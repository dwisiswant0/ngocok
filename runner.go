package main

import (
	"context"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"golang.ngrok.com/ngrok"
)

type runner struct {
	opt    *options
	ID     string
	tunnel ngrok.Tunnel

	*os.File
	*http.Server
}

func (r *runner) start(ctx context.Context) error {
	if r.opt.Output != "" {
		file, err := os.OpenFile(
			r.opt.Output,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)

		if err != nil {
			return err
		}

		r.File = file
		httplog.Colors = false
		httplog.SetOutput(file)
	}

	if !r.opt.Unstrip {
		httplog.SkipHeader(skipHeaders)
	}

	tunnel, err := r.getTunnel()
	if err != nil {
		return err
	}

	authtoken, err := r.getAuthtokenOpt()
	if err != nil {
		return err
	}

	listener, err := ngrok.Listen(
		ctx, tunnel, authtoken,
		ngrok.WithDisconnectHandler(disconnected),
		ngrok.WithConnectHandler(connected),
	)

	if err != nil {
		return err
	}

	r.tunnel = listener

	url, err := parseURL(listener.URL())
	if err != nil {
		return err
	}

	log.InfoContext(ctx,
		"server started",
		"url", url,
		"id", listener.ID(),
	)

	r.ID = listener.ID()
	handler := http.HandlerFunc(r.handler)
	r.Server = &http.Server{Handler: handler}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func(ctx context.Context) {
		if err := r.notify(sig); err != nil {
			log.ErrorContext(ctx, err.Error())
			os.Exit(1)
		}
	}(ctx)

	return r.Serve(listener)
}

func (r *runner) end(sig os.Signal) error {
	log.Warn("gracefully shutdown", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// log.Debug("closing tunnel")
	err := r.tunnel.CloseWithContext(ctx)
	if err != nil {
		return err
	}

	// log.Debug("shutting down server")
	err = r.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
