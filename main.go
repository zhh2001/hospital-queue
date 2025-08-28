package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"hospital-queue/handlers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", handlers.MainHandler)
	router.GET("/index", handlers.IndexHandler)

	group := router.Group("/queue")
	{
		group.GET("/", handlers.GetAllQueuesHandler)
		group.POST("/new", handlers.CreateQueueHandler)
		group.POST("/call", handlers.CallQueueHandler)
	}

	srv := &http.Server{
		Addr:    ":3216",
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
