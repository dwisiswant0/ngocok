package main

import (
	"flag"
	"fmt"
	// "io"
	"os"
	"time"

	"log/slog"

	"github.com/henvic/httpretty"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

func init() {
	log = slog.New(
		tint.NewHandler(colorable.NewColorableStderr(),
			&tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.DateTime,
			},
		),
	)

	httplog = &httpretty.Logger{
		RequestHeader: true,
		RequestBody:   true,
		Colors:        true,
		SkipSanitize:  true,
	}

	opt = new(options)

	flag.StringVar(&opt.Authtoken, "t", "", "")
	flag.StringVar(&opt.Authtoken, "token", "", "")

	flag.StringVar(&opt.Endpoint, "e", "http", "")
	flag.StringVar(&opt.Endpoint, "endpoint", "http", "")

	flag.BoolVar(&opt.Unstrip, "unstrip", false, "")

	flag.StringVar(&opt.Output, "o", "", "")
	flag.StringVar(&opt.Output, "output", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s%s\n", banner, usage)
	}
	flag.Parse()

	fmt.Fprintf(os.Stderr, "%s\n", banner)
}
