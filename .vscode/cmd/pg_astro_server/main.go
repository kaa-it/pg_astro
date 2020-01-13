package main

import (
	"github.com/kaa-it/pg_astro"
	"time"
)

const (
	maxReadBytes          = 1000
	idleControllerTimeout = 5 // Minutes
	port                  = "9896"
)

func main() {

	srv := pg_astro.Server{
		Addr:         ":" + port,
		IdleTimeout:  idleControllerTimeout * time.Minute,
		MaxReadBytes: maxReadBytes,
	}

	srv.ListenAndServe()
}
