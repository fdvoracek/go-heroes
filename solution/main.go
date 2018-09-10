package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type helloServer struct {
	server http.Server
}

func (hs *helloServer) start() {
	go func() {
		if err := hs.server.ListenAndServe(); err != nil {
			panic(err)
		}
		defer hs.server.Close()
	}()
}

func main() {
	fmt.Println("Starting server ...")

	hello := &helloServer{
		server: http.Server{Addr: ":8080"},
	}

	go hello.start()
	fmt.Println("Server started")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

}
