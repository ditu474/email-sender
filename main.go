package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ditu474/email-sender/handlers"
	"github.com/ditu474/email-sender/middlewares"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "SendEmail-API ", log.LstdFlags)
}

func main() {
	h := handlers.NewSendEmail()

	sm := http.NewServeMux()
	sm.Handle("/sendEmail", middlewares.CORSMiddleware(h))

	s := http.Server{
		Addr:         ":8000",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("The server fails starting: %+s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)
	sig := <-c
	l.Println("Got signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		l.Fatalf("The server fails shuting down: %+s", err)
	}
	l.Println("Server exited properly")
}
