package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var help bool
var port int

func main() {
	flag.BoolVar(&help, "help", false, "Displays contextual help")
	flag.IntVar(&port, "port", 3000, "The port to run the server on")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	err := ServeStatic(port)
	if err != nil {
		log.Fatalln(err)
	}
}

func ServeStatic(port int) error {
	fmt.Printf("Serving files on localhost: %d", port)

	host := fmt.Sprintf("localhost:%d", port)

	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
