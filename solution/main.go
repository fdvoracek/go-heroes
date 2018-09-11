package main

import (
	"fmt"
	"github.com/CoufalJa/go-workshop/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting server ...")
	helloServer := server.NewHelloServer()
	go helloServer.Start()
	fmt.Println("Server started")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
