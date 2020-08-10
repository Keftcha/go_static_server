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
var baseURL string = "/"

func serveFiles(w http.ResponseWriter, r *http.Request) {
	// Allow CORS request
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Log the request
	fmt.Println(
		fmt.Sprintf(
			"%s - %s %s",
			time.Now().Format("[2006-01-02 15:04:05]"),
			r.Host,
			r.URL,
		),
	)

	// Remove base url and serve the content
	http.StripPrefix(
		baseURL,
		http.FileServer(http.Dir(dir)),
	).ServeHTTP(w, r)
}

func init() {
	// Check is the user gave the address os set a default one
	flag.StringVar(&host, "host", "0.0.0.0", "The host where the server listen")

	// Check if the user gave the port or set a default port
	flag.IntVar(&port, "port", 8080, "The port where ther server listen")
	flag.IntVar(&port, "p", 8080, "The port where ther server listen")

	// Check is the user gave a directory to serve
	flag.StringVar(&dir, "dir", "./", "The directory to serve")
	flag.StringVar(&dir, "d", "./", "The directory to serve")

	// Check if the user have a base url to strip
	flag.StringVar(&baseURL, "base-url", "/", "The base url")

	flag.Parse()

	// Check if the given directory to serve exist
	if _, err := os.Stat(dir); err != nil {
		panic(err)
	}
}

func main() {
	// Serve all as static file
	http.HandleFunc("/", serveFiles)

	// Start the server
	fmt.Println(fmt.Sprintf("Serving the directory %s on http://%s:%d%s ", dir, host, port, baseURL))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
