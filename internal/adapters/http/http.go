package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alekseysychev/test-CRUD-auto/internal/domain/vehicle"
)

var done chan os.Signal

type HttpServer struct {
	VehicleStorage vehicle.Vehicle
}

func (hs *HttpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
}

func (hs *HttpServer) Serve(addr string) {
	http.Handle("/", hs)

	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	done = make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server Started")

	<-done
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Println("Server Exited Properly")
}
