# go-boiler - Golang Boilerplate

go-boiler is a boilerplate project for kickstarting your Golang applications. It provides a basic project structure with essential components and configurations to help you get started quickly.

## Project Structure

```sh
├── api
│ ├── routes.go
│ ├── server.go
│ └── v1
│ ├── handlers.go
│ └── routes.go
├── cmd
│ └── root_cmd.go
├── conf
│ ├── config.go
│ ├── db.go
│ ├── fluent.go
│ ├── logging.go
│ └── redis.go
├── config
│ └── config.json
├── Dockerfile
├── go-boiler.yml
├── go.mod
├── go.sum
├── internal
│ └── internal.go
├── main.go
├── readme.md
└── utils
└── utils.go
``````

- `api`: Contains routing and server setup.
- `cmd`: Command-line application entry point.
- `conf`: Configuration management.
- `config`: Configuration files.
- `Dockerfile`: Docker container configuration.
- `go-boiler.yml`: Project-specific configuration (if any).
- `internal`: Internal package (not for external use).
- `main.go`: Main application entry point.
- `utils`: Utility functions.

## Getting Started

1. Clone the repository:

   ```sh
   git clone https://github.com/rohanchavan1918/go-boiler.git
   cd go-boiler

### Build and run your application

```sh
go build -o myapp
./myapp
```

Access your application at http://localhost:8080.

### Configuration

- Modify configuration settings in the config/config.json file
- Configuration is loaded using Viper

### Dependencies

This project uses the following external dependencies (as listed in go.mod):

- [logrus](https://github.com/sirupsen/logrus): Structured logger for Go.
- [cobra](https://github.com/spf13/cobra): CLI library for Go.
- [viper](https://github.com/spf13/viper): Configuration management for Go.
- [lumberjack](https://github.com/natefinch/lumberjack.v2): Rolling log files for Go applications.

Feel free to extend and customize this project to suit your specific needs.

### License

This project is licensed under the MIT License - see the LICENSE file for details.