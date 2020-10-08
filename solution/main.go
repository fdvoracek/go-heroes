package main

import (
	"flag"
	"fmt"
	"github.com/CoufalJa/go-workshop/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

var name string


func main() {
	flag.Parse()

	fmt.Println("Starting server ...")
	helloServer := server.NewHelloServer(name)
	go helloServer.Start()
	fmt.Println("Server started")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	fmt.Println("Stopping server")
}
