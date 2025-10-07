# Cobra vs pflag Comparison

This directory contains examples demonstrating the differences between `pflag` and `cobra` when working with Viper.

## Files

- `main.go` - Original example using **pflag** directly
- `cobra_main.go` - Example using **Cobra** framework
- `config.yaml` - Sample configuration file

## Key Differences

### pflag (main.go)
```go
// Direct flag definition
pflag.String("config", "", "config file")
pflag.String("app.name", "", "application name")
pflag.Parse()

// Direct viper binding
viper.BindPFlags(pflag.CommandLine)
```

**Pros:**
- Simple and lightweight
- Direct control over flag parsing
- Good for single-command applications

**Cons:**
- Manual help text generation
- No subcommand support
- Limited structure for complex CLIs

### Cobra (cobra_main.go)
```go
// Command structure with automatic help
var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Application description",
    Run: func(cmd *cobra.Command, args []string) {
        // Command logic
    },
}

// Subcommands
rootCmd.AddCommand(configCmd)
configCmd.AddCommand(showConfigCmd)
```

**Pros:**
- Automatic help generation
- Subcommand support (`app config show`)
- Professional CLI structure
- Built-in shell completion
- Better organization for complex tools

**Cons:**
- More setup code required
- Heavier dependency
- Overkill for simple utilities

## When to Use Each

### Use pflag when:
- Building simple, single-purpose tools
- You need lightweight dependency
- Single command with just flags
- Learning/prototyping

### Use Cobra when:
- Building professional CLI tools
- Need subcommands (like `git add`, `docker run`)
- Want automatic help/usage generation
- Building complex applications with multiple functions
- Used by tools like: Docker, Kubernetes, GitHub CLI, Hugo

## Running Examples

### pflag version (main.go):
```bash
# Basic usage
go run main.go

# With flags
go run main.go --app.debug --server.port=8080

# With environment variables
VIPEREXAMPLE_APP_NAME="pflag App" go run main.go

# Help (basic)
go run main.go --help  # Shows basic usage
```

### Cobra version (cobra_main.go):
```bash
# Basic usage
go run cobra_main.go

# Subcommands
go run cobra_main.go config show
go run cobra_main.go config validate

# With flags
go run cobra_main.go --app.debug --server.port=8080 config show

# With environment variables
VIPEREXAMPLE_APP_NAME="Cobra App" go run cobra_main.go config show

# Professional help system
go run cobra_main.go --help           # Shows main help
go run cobra_main.go config --help    # Shows subcommand help
go run cobra_main.go completion bash  # Generate shell completion
```

## Real-world Examples

### Tools using pflag:
- Simple utilities
- Single-purpose command line tools
- Quick scripts and prototypes

### Tools using Cobra:
- **Docker**: `docker run`, `docker build`, `docker ps`
- **Kubernetes**: `kubectl get`, `kubectl apply`, `kubectl delete`
- **GitHub CLI**: `gh repo create`, `gh pr list`
- **Hugo**: `hugo new`, `hugo server`
- **Helm**: `helm install`, `helm upgrade`

## Integration with Viper

Both work well with Viper, but Cobra provides:
- Better flag binding mechanisms
- Persistent flags across subcommands
- Automatic configuration initialization
- More structured approach to configuration management