package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ota-server/internal/config"
	"syscall"
	"time"
)

func main() {
	config.LoadEnv()
	app := gin.Default()

	listenAndServe(app, true)
}

func listenAndServe(app *gin.Engine, gracefully bool) {
	var err error
	if gracefully {
		server := &http.Server{
			Addr:    ":8080",
			Handler: app.Handler(),
		}
		go func() {
			if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}()
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		<-exit
		log.Println("Shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err = server.Shutdown(ctx); err != nil {
			log.Fatalln("Fail to shutdown server with timeout:", err)
		}
		select {
		case <-ctx.Done():
			log.Println("Server shutdown successfully!")
		default:
			log.Println("Server shutdown successfully!")
		}
	} else {
		if err = app.Run(":8080"); err != nil {
			panic(err)
		}
	}
}
