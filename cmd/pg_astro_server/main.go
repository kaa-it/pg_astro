package main

import (
	"github.com/kaa-it/pg_astro"
	"sync"
	"time"
)

const (
	maxReadBytes          = 5000
	idleTimeout           = 5 // Seconds
	idleControllerTimeout = 5 // Minutes
	port                  = "9896"
)

func main() {

	srv := pg_astro.Server{
		Addr:         ":" + port,
		IdleTimeout:  idleControllerTimeout * time.Second,
		MaxReadBytes: maxReadBytes,
	}

	var wg sync.WaitGroup

	pg_astro.InitSignals(&wg, &srv)

	srv.ListenAndServe()

	wg.Wait()
}
