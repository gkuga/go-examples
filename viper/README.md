# Viper Configuration Management Example

This example demonstrates how to use [Viper](https://github.com/spf13/viper) for configuration management in Go applications.

## Features Demonstrated

This sample application showcases the following Viper capabilities:

1. **Multiple Configuration Sources** with priority order:
   - Command line flags (highest priority)
   - Environment variables
   - Configuration files (YAML/JSON)
   - Default values (lowest priority)

2. **Configuration File Formats**:
   - YAML configuration (`config.yaml`)
   - JSON configuration (`config.json`)
   - Automatic file type detection

3. **Environment Variables**:
   - Automatic environment variable binding
   - Environment variable prefix (`VIPEREXAMPLE_`)
   - Key replacement (dots to underscores)

4. **Data Types**:
   - Strings, integers, booleans
   - Nested configuration objects
   - Map structures

5. **Advanced Features**:
   - Configuration validation with `IsSet()`
   - Unmarshaling to Go structs
   - Writing configuration back to files
   - Live configuration monitoring

## Files

- `main.go` - Main application demonstrating Viper features
- `config.yaml` - Sample YAML configuration file
- `config.json` - Sample JSON configuration file (alternative)
- `go.mod` - Go module definition

## Running the Example

### Using Makefile (Recommended)

The project includes a comprehensive Makefile for easy management:

```bash
# Show all available commands
make help

# Run the pflag version
make run

# Run the cobra version
make run-cobra

# Build both executables
make build

# Run with environment variables
make run-with-env
make run-cobra-with-env

# Development tasks
make check          # Format code and run vet
make deps           # Download dependencies
make clean          # Clean build artifacts

# Run demonstrations
make demo           # Demo both versions
make demo-pflag     # Demo pflag version only
make demo-cobra     # Demo cobra version only

# Show project information
make info
```

### Manual Usage

### Manual Usage

#### Basic Usage

```bash
# Run with YAML config file
go run main.go

# Run with specific config file
go run main.go --config=config.yaml
```

### Environment Variables

Set environment variables with the `VIPEREXAMPLE_` prefix:

```bash
# Override app name and server port
export VIPEREXAMPLE_APP_NAME="My Custom App"
export VIPEREXAMPLE_SERVER_PORT=9999
export VIPEREXAMPLE_APP_DEBUG=true

go run main.go
```

### Command Line Flags

Override configuration with command line flags:

```bash
# Override specific values
go run main.go --app.debug --server.port=8888 --logging.level=debug

# Combine with environment variables
VIPEREXAMPLE_APP_NAME="CLI Example" go run main.go --server.port=3000
```

### Priority Order Example

```bash
# This demonstrates priority: CLI > ENV > Config File > Defaults
VIPEREXAMPLE_SERVER_PORT=7777 go run main.go --server.port=8888
# Result: server.port will be 8888 (CLI flag wins)
```

## Configuration Structure

The application uses the following configuration structure:

```yaml
app:
  name: "Application Name"
  version: "1.0.0"
  debug: false

server:
  host: "localhost"
  port: 8080
  timeout: 30

database:
  driver: "postgresql"
  host: "localhost"
  port: 5432
  name: "myapp"
  user: "admin"
  password: "secret123"
  ssl_mode: "disable"

logging:
  level: "info"
  file: "app.log"
  max_size: 100
  max_backups: 3
  max_age: 28

features:
  enable_metrics: true
  enable_tracing: false
  max_connections: 100
```

## Environment Variable Mapping

Environment variables use the prefix `VIPEREXAMPLE_` and replace dots with underscores:

- `app.name` → `VIPEREXAMPLE_APP_NAME`
- `server.port` → `VIPEREXAMPLE_SERVER_PORT`
- `database.ssl_mode` → `VIPEREXAMPLE_DATABASE_SSL_MODE`

## Key Viper Concepts

### Configuration Sources Priority

1. **Command Line Flags** (highest)
2. **Environment Variables**
3. **Configuration Files**
4. **Default Values** (lowest)

### Key Methods Used

- `viper.SetDefault()` - Set default values
- `viper.SetConfigName()` - Set config file name
- `viper.AddConfigPath()` - Add search paths
- `viper.SetEnvPrefix()` - Set environment variable prefix
- `viper.AutomaticEnv()` - Enable automatic env var binding
- `viper.ReadInConfig()` - Read configuration file
- `viper.Get*()` - Get values with type safety
- `viper.Unmarshal()` - Unmarshal to struct
- `viper.IsSet()` - Check if value is explicitly set

### Struct Mapping

The example shows how to unmarshal configuration into Go structs using the `mapstructure` tags:

```go
type Config struct {
    App struct {
        Name    string `mapstructure:"name"`
        Version string `mapstructure:"version"`
        Debug   bool   `mapstructure:"debug"`
    } `mapstructure:"app"`
    // ... more fields
}
```

## Dependencies

- `github.com/spf13/viper` - Configuration management
- `github.com/spf13/pflag` - POSIX/GNU-style command line flags

## Learning Resources

- [Viper Documentation](https://github.com/spf13/viper)
- [Viper Examples](https://github.com/spf13/viper/tree/master/_examples)
- [Configuration Management Best Practices](https://12factor.net/config)