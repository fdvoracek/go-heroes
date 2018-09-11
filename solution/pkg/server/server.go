package server

import (
	"fmt"
	"net/http"
)

type HelloServer interface {
	Start()
}

type helloServer struct {
	server http.Server
}

func (hs *helloServer) Start() {
	go func() {
		if err := hs.server.ListenAndServe(); err != nil {
			panic(err)
		}
		defer hs.server.Close()
	}()
}

func (hs *helloServer) handleHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Stranger")
}

func (hs *helloServer) handleGoodbye(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Goodbye Stranger")
}

func NewHelloServer() HelloServer {
	hello := &helloServer{
		server: http.Server{Addr: ":8080"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello.handleHello)
	mux.HandleFunc("/goodbye", hello.handleGoodbye)
	hello.server.Handler = mux

	return hello
}
