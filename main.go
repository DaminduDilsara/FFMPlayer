package main

import (
	"AdvancedNetwork/connections"
	"AdvancedNetwork/pkg/apis"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"path/filepath"
)

var (
	path = flag.String("path", "./segments", "path to the folder to serve. Defaults to the current folder")
	port = flag.String("port", "8080", "port to serve on. Defaults to 8080")
)

func main() {
	fmt.Println("Starting UI Generator Application")
	connections.ConnectMongo()

	e := echo.New()
	api.EchoManager(e)

	go createFileServer()

	go e.Logger.Fatal(e.Start(":4000"))

}

func Serve(dirname string, port string) error {
	fs := http.FileServer(http.Dir(dirname))
	http.Handle("/", fs)

	return http.ListenAndServe(":"+port, nil)
}

func createFileServer() {
	flag.Parse()
	dirname, err := filepath.Abs(*path)
	if err != nil {
		log.Fatalf("Could not get absolute path to directory: %s: %s", dirname, err.Error())
	}
	log.Printf("Serving %s on port %s", dirname, *port)
	err = Serve(dirname, *port)
	if err != nil {
		log.Fatalf("Could not serve directory: %s: %s", dirname, err.Error())
	}
}
