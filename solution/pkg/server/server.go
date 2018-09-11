package server

import (
	"encoding/json"
	"fmt"
	"github.com/CoufalJa/go-workshop/pkg/model"
	"net/http"
)

type HelloServer interface {
	Start()
}

type helloServer struct {
	name   string
	server http.Server
}

func (hs *helloServer) Start() {
	if err := hs.server.ListenAndServe(); err != nil {
		panic(err)
	}
	defer hs.server.Close()
}

func (hs *helloServer) handleHello(writer http.ResponseWriter, request *http.Request) {
	bytes, err := json.Marshal(model.Saying{Name: hs.name})
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, string(bytes))
}

func NewHelloServer(name string) HelloServer {
	hello := &helloServer{
		name:   name,
		server: http.Server{Addr: ":8080"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello.handleHello)
	hello.server.Handler = mux

	return hello
}
