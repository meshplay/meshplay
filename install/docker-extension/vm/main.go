package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/docker/meshplay-extension/vm/proxy"
	"github.com/sirupsen/logrus"
)

var (
	MeshplayServerHost = "http://host.docker.internal:9081"
	Port              = "7877"
)

func main() {

	serveAt := fmt.Sprintf("0.0.0.0:%s", Port)
	ln, err := net.Listen("tcp", serveAt)
	if err != nil {
		log.Fatal(err)
	}

	handler := &proxy.Proxy{}
	logrus.New().Infof("Starting listening on %s \n", serveAt)
	if err := http.Serve(ln, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
