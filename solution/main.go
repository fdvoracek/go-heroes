package solution

import (
	"flag"
	"fmt"
	"github.com/fdvoracek/go-heroes/solution/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()

	fmt.Println("Starting server ...")
	helloServer := server.NewHelloServer()
	go helloServer.Start()
	fmt.Println("Server started")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	fmt.Println("Stopping server")
}
