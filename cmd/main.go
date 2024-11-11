package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ota-server/api"
	"ota-server/internal/config"
	"syscall"
	"time"
)

func main() {
	config.LoadEnv()
	app := gin.Default()

	// Attach endpoints
	api.APIEntryRouter(app)

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
		// Wait until 30 seconds for server shutdown - For long term connections
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err = server.Shutdown(ctx); err != nil {
			log.Fatalln("Fail to shutdown server with timeout:", err)
		}
		log.Println("Server shutdown gracefully!")
	} else {
		if err = app.Run(":8080"); err != nil {
			panic(err)
		}
	}
}
