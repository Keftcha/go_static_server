# Go Static Server (gss)

Little web static server to serve static files.

## Usage

### Compile the gss

1. Download, clone or `go get` this project
2. Compile it with `go install` or `go build` (not needed if you use `go get`)
3. Start the server

### Compile the gss inside docker

If you don't have go installed (and you don't want to install it), you can
build the gss inside a docker container.

1. Download or clone this project
2. Compile inside a docker container with the command:  
`docker run -v "$PWD":/go/src/gss -w /go/src/gss golang go build`
3. Start the server

## Command Line options

- `--help` or `-h` flag will list available options
- `--host` flag define the host address the server listen (default is "0.0.0.0")
- `--port` or `-p` flag define the listen port number (default is 8080)
- `--dir` or `-d` flag define the directory to serve (default is the current directory aka. "./")
- `--base-url` flag define the base url of the server (default is "/").
