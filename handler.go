package main

import (
	"context"
	"os"
	"strings"

	"net/http"

	"golang.ngrok.com/ngrok"
)

func (r *runner) handler(w http.ResponseWriter, req *http.Request) {
	r.logreq(req)

	for key, value := range responseHeaders {
		w.Header().Set(key, value)
	}

	_, err := w.Write([]byte(r.ID))
	if err != nil {
		log.Warn("cannot write response", "err", err)
	}
}

func (r *runner) notify(sigCh chan os.Signal) error {
	sig := <-sigCh

	return r.end(sig)
}

func connected(ctx context.Context, sess ngrok.Session) {
	log.InfoContext(ctx, sessOK)
}

func disconnected(ctx context.Context, sess ngrok.Session, err error) {
	var errs, msgs []string

	if strings.Contains(err.Error(), ":") {
		errs = strings.Split(err.Error(), "\n")
		msgs = strings.Split(errs[0], ":")
		for i, msg := range msgs {
			msgs[i] = strings.TrimSpace(msg)
		}

		log.ErrorContext(ctx, msgs[0], "err", msgs[1])

		// NOTE(dwisiswant0): yea I don't want to handle with `ngrok.Error`
		// 'cuz that'll produces switch..case [CHANGE MY MIND]
		if strings.Contains(err.Error(), authFailed) {
			log.InfoContext(ctx, infGetAuthtokenMsg, "ref", infGetAuthtokenRef)
			os.Exit(2)
		}
	}

	log.ErrorContext(ctx, err.Error())
}
