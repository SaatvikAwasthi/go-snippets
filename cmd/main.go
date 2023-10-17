package main

import (
	"flag"
	"go-snips/server"
	"log"
	"strconv"
)

const (
	port            = "port"
	portUsage       = "use to specify port"
	defaultPort int = 8080
)

var httpPort string

func init() {
	args := flag.Int(port, defaultPort, portUsage)
	log.Printf("Port: %v", *args)
	httpPort = strconv.Itoa(*args)
}

func main() {
	server.Server(httpPort)
}
