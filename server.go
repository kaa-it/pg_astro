package pg_astro

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

var printMx sync.Mutex

type Server struct {
	Addr         string
	IdleTimeout  time.Duration
	MaxReadBytes int64

	listener   net.Listener
	conns      map[*conn]struct{}
	mu         sync.Mutex
	inShutdown bool
	done       chan struct{}
}

func (srv *Server) ListenAndServe() error {
	srv.done = make(chan struct{})
	addr := srv.Addr
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("starting server on %v\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	srv.listener = listener
	for {
		// should be guarded by mu
		if srv.inShutdown {
			break
		}
		newConn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		log.Printf("accepted connection from %v", newConn.RemoteAddr())
		conn := &conn{
			Conn:          newConn,
			IdleTimeout:   srv.IdleTimeout,
			MaxReadBuffer: srv.MaxReadBytes,
		}
		srv.trackConn(conn)
		conn.SetDeadline(time.Now().Add(conn.IdleTimeout))
		go srv.handle(conn)
	}
	return nil
}

func (srv *Server) trackConn(c *conn) {
	defer srv.mu.Unlock()
	srv.mu.Lock()
	if srv.conns == nil {
		srv.conns = make(map[*conn]struct{})
	}
	srv.conns[c] = struct{}{}
}

func (srv *Server) handle(conn *conn) error {
	defer func() {
		log.Printf("closing connection from %v", conn.RemoteAddr())
		conn.Close()
		srv.deleteConn(conn)
	}()

	w := bufio.NewWriter(conn)

	var wg sync.WaitGroup
	wg.Add(1)

	var lastTime = time.Now()

	go func() {
		var b [4096]byte
		buffer := make([]byte, 0)

		for {

			len, err := conn.Read(b[:])
			if err != nil {
				if errors.Is(err, syscall.ETIMEDOUT) {
					if srv.inShutdown {
						break
					}
					if time.Since(lastTime) > 5*time.Minute {
						break
					} else {
						lastTime = time.Now()
						continue
					}
				}
				// TODO: Log error
				break
			}

			extractPackage(b[:len], buffer)

		}
		wg.Done()
	}()

	w.WriteString(`{"node":"0xFFFFFFFF","list":0}`)
	w.Flush()

	wg.Wait()

	return nil
}

func extractPackage(message []byte, buffer []byte) {
	s := bytes.Trim(message, "\n \t")
	endOfPackage := bytes.Index(s, []byte("}"))

	if endOfPackage == -1 {
		// String doesn't contain '}'
		buffer = append(buffer, s...)
	} else {
		// String contains '}'
		startOfPackage := bytes.LastIndex(buffer, []byte("{"))

		if startOfPackage != -1 {
			// Start of package in the buffer
			pkg := buffer[startOfPackage:]
			buffer = buffer[:startOfPackage]
			pkg = append(pkg, s[:endOfPackage+1]...)
			//parsePackage(pkg)
			fmt.Println(string(pkg))
			extractPackage(s[endOfPackage+1:], buffer)
		} else {
			// Start of package doesn't exists in buffer
			startOfPackage = bytes.Index(s, []byte("{"))

			if startOfPackage == -1 {
				fmt.Println("Wrong package received")
			} else {
				//parsePackage(s[startOfPackage, endOfPackage + 1])
				fmt.Println(string(s[startOfPackage : endOfPackage+1]))
				extractPackage(s[endOfPackage+1:], buffer)
			}
		}
	}
}

func (srv *Server) deleteConn(conn *conn) {
	defer srv.mu.Unlock()
	srv.mu.Lock()
	delete(srv.conns, conn)
}

func (srv *Server) Shutdown() {
	// should be guarded by mu
	srv.inShutdown = true
	log.Println("shutting down...")
	srv.listener.Close()
	close(srv.done)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Printf("waiting on %v connections", len(srv.conns))
		}
		if len(srv.conns) == 0 {
			return
		}
	}
}
