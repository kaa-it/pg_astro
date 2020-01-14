package main

import (
	"github.com/kaa-it/pg_astro"
	"sync"
	"time"
)

const (
	maxReadBytes          = 5000
	idleTimeout           = 5 * time.Second
	idleControllerTimeout = 5 * time.Minute
	port                  = "9896"
)

func main() {

	srv := pg_astro.Server{
		Addr:                  ":" + port,
		IdleTimeout:           idleTimeout,
		IdleControllerTimeout: idleControllerTimeout,
		MaxReadBytes:          maxReadBytes,
	}

	var wg sync.WaitGroup

	pg_astro.InitSignals(&wg, &srv)

	srv.ListenAndServe()

	wg.Wait()
}
