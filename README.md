# Go Static Server (gss)

Little web static server to serve static files.

## Usage

### Compile the gss

1. Download or clone this project
2. Compile it:
    - With `make build` to get a bianry named `gss`.
    - With `make install` to install the executable in the `$GOBIN` directory
    - With `make-in-ctn` to build it in a container and get a binary named `gss`
3. Start the server

## Command Line options

- `--help` or `-h` flag will list available options
- `--host` flag define the host address the server listen (default is "0.0.0.0")
- `--port` or `-p` flag define the listen port number (default is 8080)
- `--dir` or `-d` flag define the directory to serve (default is the current directory aka. "./")
- `--base-url` flag define the base url of the server (default is "/").
