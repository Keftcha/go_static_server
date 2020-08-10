# Go Static Server (gss)

Little web static server to serve static files.

## Usage

1. Download, clone or `go get` this project
2. Compile it with `go install` or `go build` (not needed if you use `go get`)
3. Start the server

## Command Line options

- `--help` or `-h` flag will list available options
- `--host` flag define the host address the server listen (default is "0.0.0.0")
- `--port` or `-p` flag define the listen port number (default is 8080)
- `--dir` or `-d` flag define the directory to serve (default is the current directory aka. "./")
- `--base-url` flag define the base url of the server (default is "/").
