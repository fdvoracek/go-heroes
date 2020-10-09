package server

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fdvoracek/go-heroes/solution/pkg/db"
	"github.com/fdvoracek/go-heroes/solution/pkg/model"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type HelloServer interface {
	Start()
}

type helloServer struct {
	server         http.Server
	memcacheClient db.MemcacheClient
}

func (hs *helloServer) Start() {
	if err := hs.server.ListenAndServe(); err != nil {
		panic(err)
	}
	defer hs.server.Close()
}

func hashToSha256(data string) []byte {

	hash := sha256.New()
	hash.Write([]byte(data))

	return hash.Sum(nil)
}

func (hs *helloServer) handleFilterWithChain(writer http.ResponseWriter, request *http.Request) {

	var filterRequest model.Request
	json.NewDecoder(request.Body).Decode(&filterRequest)

	hashedDomain := hashToSha256(filterRequest.Domain)

	chain := make(chan model.SecurityDefinition)

	var expectedArrayLength = 3
	for i := 0; i< expectedArrayLength; i++ {
		go hs.memcacheClient.Get(hashedDomain, filterRequest.Domain)
	}

	responses := make([]model.SecurityDefinition, expectedArrayLength)

	for i, _ := range responses {
		responses[i] = <-chain
	}

	bytes, err := json.Marshal(responses)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(writer, string(bytes))
}

func (hs *helloServer) handleFilter(writer http.ResponseWriter, request *http.Request) {
	var filterRequest model.Request
	json.NewDecoder(request.Body).Decode(&filterRequest)
	hashedDomain := hashToSha256(filterRequest.Domain)
	const expectedArrayLength = 3
	chain := make(chan model.SecurityDefinition, expectedArrayLength)
	defer close(chain)
	var wg sync.WaitGroup
	for i := 0; i< expectedArrayLength; i++ {
		wg.Add(1)
		go func() {
			chain <- hs.memcacheClient.Get(hashedDomain, filterRequest.Domain)
			wg.Done()
		}()
	}
	wg.Wait()
	responses := [expectedArrayLength]model.SecurityDefinition{}
	for i, _ := range responses {
		responses[i] = <-chain
	}
	
	err := json.NewEncoder(writer).Encode(responses)
	if err != nil {
		panic(err)
	}
}

func NewHelloServer() HelloServer {
	//var requestTimeout = 150 * time.Millisecond
	var requestTimeout = 10 * time.Second

	hello := &helloServer{
		server:         http.Server{Addr: ":8080"},
		memcacheClient: db.NewMemcacheClient("127.0.0.1:11211", requestTimeout),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/performancetest/security-domain", hello.handleFilter)
	// Register pprof handlers
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	hello.server.Handler = mux

	return hello
}
