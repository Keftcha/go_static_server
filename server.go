package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var host string = "0.0.0.0"
var port int = 8080
var dir string = "./"
var baseURL string = "/"

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
	// Create the server
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: http.HandlerFunc(serveFiles),
	}

	// Gracefully shutdown the server on SIGINT or SIGTERM
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		<-sc
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	// Start the server
	fmt.Println(fmt.Sprintf("Serving the directory %s on http://%s:%d%s ", dir, host, port, baseURL))
	// Ignore the ErrServerClosed because we may have close it
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	// Allow CORS request
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Log the request
	fmt.Println(
		fmt.Sprintf(
			"%s - %s - %s %s %s",
			time.Now().Format("[2006-01-02 15:04:05]"),
			r.RemoteAddr,
			r.Host,
			r.Method,
			r.URL,
		),
	)

	// Remove base url and serve the content
	http.StripPrefix(
		baseURL,
		http.FileServer(http.Dir(dir)),
	).ServeHTTP(w, r)
}
