package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var host string = "0.0.0.0"
var port int64 = 8080
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
	args := os.Args[1:]

	// Check is the user gave the address os set a default one
	if _, idx, okHost := seekable(args).contain("--host"); okHost {
		host = args[idx+1]
	}

	// Check if the user gave the port or set a default port
	if _, idx, okPort := seekable(args).contain("--port"); okPort { // && (_, idx, okP := seekable(args).contain("-p"); okPort && okP) {
		givenPort, err := strconv.ParseInt(args[idx+1], 10, 16)
		if err != nil {
			panic(err)
		}
		port = givenPort
	}

	// Check is the user gave a folder to serve and check if the directory exist
	if _, idx, okDir := seekable(args).contain("--dir"); okDir { // && _, idx, okD := seekable(args).contain("-d") {
		givenDir := args[idx+1]
		if _, err := os.Stat(givenDir); err != nil {
			panic(err)
		}
		dir = givenDir
	}
}

func main() {

	// Serve all as static file
	http.Handle("/", serveFiles{})

	// Start the server
	fmt.Println(fmt.Sprintf("Serving the directory %s on %s:%d ", dir, host, port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
