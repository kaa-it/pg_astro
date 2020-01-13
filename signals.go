package pg_astro

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func InitSignals(wg *sync.WaitGroup, srv *Server) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fallthrough
		case syscall.SIGTERM:
			signal.Stop(signalChannel)
			close(signalChannel)
			srv.Shutdown()
			fmt.Println("Signals thread done....")
			wg.Done()
			return
		}
	}()
}
