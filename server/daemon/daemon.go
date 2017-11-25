package daemon

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
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

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	handlerWithCORS := cors.Default().Handler(loggedRouter)
	http.Handle("/", handlerWithCORS)

	server.Serve(l)
}
