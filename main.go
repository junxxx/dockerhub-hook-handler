package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/junxxx/dockerhub-webhook/internal"
)

var port string

func main() {
	flag.StringVar(&port, "port", "29403", "server port")
	http.HandleFunc("/", internal.HookHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
