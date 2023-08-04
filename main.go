package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/psi-anamnese/psi-anamnese-be/infra/http/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	routes.NewRoutes(r).Setup()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
		log.Println("Server exiting")
	}

}
