package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var host string = "0.0.0.0"
var port int = 8080
var dir string = "./"

type seekable []string
type serveFiles struct{}

func (slice seekable) contain(toFind string) (string, int, bool) {
	// Given a slice and a thing to find in that slice
	// we return:
	//    - the value of the finded element
	//    - the index of the finded element
	//    - a boolean to know if the value is present

	for idx, elem := range slice {
		if elem == toFind {
			return elem, idx, true
		}
	}
	return "", 0, false
}

func (h serveFiles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(
		fmt.Sprintf(
			"%s - %s %s",
			time.Now().Format("[2006-01-02 15:04:05]"),
			r.Host,
			r.URL,
		),
	)
	http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
}

func init() {
	// Check is the user gave the address os set a default one
	flag.StringVar(&host, "host", "0.0.0.0", "The host where the server listen")

	// Check if the user gave the port or set a default port
	flag.IntVar(&port, "port", 8080, "The port chere ther server listen")
	flag.IntVar(&port, "p", 8080, "The port chere ther server listen")

	// Check is the user gave a directory to serve
	flag.StringVar(&dir, "dir", "./", "The directory to serve")
	flag.StringVar(&dir, "d", "./", "The directory to serve")

	flag.Parse()

	// Check if the given directory to serve exist
	if _, err := os.Stat(dir); err != nil {
		panic(err)
	}
}

func main() {

	// Serve all as static file
	http.Handle("/", serveFiles{})

	// Start the server
	fmt.Println(fmt.Sprintf("Serving the directory %s on http://%s:%d ", dir, host, port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
