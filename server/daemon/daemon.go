package daemon

import (
	"log"
	"net"
	"net/http"
	"time"
)

// Config --
type Config struct {
	Listenspec string
}

// StartServer --
func StartServer(cfg Config) {
	log.Printf("Listening on %v\n", cfg.Listenspec)
	l, err := net.Listen("tcp", cfg.Listenspec)
	if err != nil {
		log.Fatalf("Error creating Listener: %v", err)
	}

	server := &http.Server{
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}
	router := initRoutes()
	http.Handle("/", router)
	server.Serve(l)
}
