package main

import (
	"fmt"
	. "github.com/pivotal-sydney/whiteboardbot/app"
	. "github.com/pivotal-sydney/whiteboardbot/http"
	"os"
	"os/signal"
	"syscall"
)

var redisConnectionPool = NewPool()

func init() {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt)
	signal.Notify(shutdownChannel, syscall.SIGTERM)
	go func() {
		<-shutdownChannel
		cleanup()
		os.Exit(1)
	}()
}

func cleanup() {
	if redisConnectionPool != nil {
		fmt.Println("Closing Redis connection pool")
		redisConnectionPool.Close()
	}
}

func main() {
	store := RealStore{Pool: redisConnectionPool}
	WhiteboardHttpServer{Store: &store}.Run()
}
