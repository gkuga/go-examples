package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "viper-cobra-example",
	Short: "A sample application demonstrating Viper with Cobra",
	Long: `This application demonstrates how to use Viper for configuration management
with Cobra for CLI structure. It shows how to handle configuration files,
environment variables, and command-line flags in a structured way.

Examples:
  viper-cobra-example                    # Run main demo
  viper-cobra-example config show        # Show configuration
  viper-cobra-example config validate    # Validate configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is called when the root command is executed
		demonstrateViperFeatures()
	},
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management commands",
	Long:  `Commands for managing application configuration`,
}

// showConfigCmd shows current configuration
var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  `Display all current configuration values from all sources (files, env vars, flags)`,
	Run: func(cmd *cobra.Command, args []string) {
		showConfiguration()
	},
}

// validateConfigCmd validates configuration
var validateConfigCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration",
	Long:  `Validate the current configuration for correctness and completeness`,
	Run: func(cmd *cobra.Command, args []string) {
		validateConfiguration()
	},
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Initialize configuration when the application starts
	cobra.OnInitialize(initCobraConfig)

	// Add persistent flags (available to all commands and subcommands)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().String("app.name", "", "application name")
	rootCmd.PersistentFlags().String("server.host", "", "server host")
	rootCmd.PersistentFlags().Int("server.port", 0, "server port")
	rootCmd.PersistentFlags().Bool("app.debug", false, "enable debug mode")
	rootCmd.PersistentFlags().String("logging.level", "", "logging level (debug, info, warn, error)")

	// Bind flags to viper (this is the key integration point)
	viper.BindPFlag("app.name", rootCmd.PersistentFlags().Lookup("app.name"))
	viper.BindPFlag("server.host", rootCmd.PersistentFlags().Lookup("server.host"))
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("server.port"))
	viper.BindPFlag("app.debug", rootCmd.PersistentFlags().Lookup("app.debug"))
	viper.BindPFlag("logging.level", rootCmd.PersistentFlags().Lookup("logging.level"))

	// Add subcommands
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(showConfigCmd)
	configCmd.AddCommand(validateConfigCmd)

	// Add local flags (only available to specific commands)
	showConfigCmd.Flags().Bool("json", false, "output in JSON format")
	showConfigCmd.Flags().Bool("yaml", false, "output in YAML format")
}

// initCobraConfig reads in config file and ENV variables
func initCobraConfig() {
	// Set default values
	setCobraDefaults()

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config file in common locations
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/appname/")
		viper.AddConfigPath("$HOME/.appname")
		viper.AddConfigPath(".")
	}

	// Environment variable configuration
	viper.SetEnvPrefix("VIPEREXAMPLE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read config file if found
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}

// setCobraDefaults sets default configuration values
func setCobraDefaults() {
	// App defaults
	viper.SetDefault("app.name", "Cobra Viper Example")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.debug", false)

	// Server defaults
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 3000)
	viper.SetDefault("server.timeout", 10)

	// Database defaults
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.name", "defaultdb")
	viper.SetDefault("database.user", "user")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.ssl_mode", "disable")

	// Logging defaults
	viper.SetDefault("logging.level", "debug")
	viper.SetDefault("logging.file", "default.log")
	viper.SetDefault("logging.max_size", 10)
	viper.SetDefault("logging.max_backups", 1)
	viper.SetDefault("logging.max_age", 7)

	// Features defaults
	viper.SetDefault("features.enable_metrics", false)
	viper.SetDefault("features.enable_tracing", false)
	viper.SetDefault("features.max_connections", 50)
}

// demonstrateViperFeatures shows various viper capabilities
func demonstrateViperFeatures() {
	fmt.Println("\n=== Cobra + Viper Configuration Demo ===")
	fmt.Println("This demonstrates Cobra's structured approach to CLI applications")

	// 1. Reading simple values
	fmt.Println("\n1. Reading Configuration Values:")
	fmt.Printf("   App Name: %s\n", viper.GetString("app.name"))
	fmt.Printf("   App Version: %s\n", viper.GetString("app.version"))
	fmt.Printf("   Debug Mode: %t\n", viper.GetBool("app.debug"))

	// 2. Reading nested values
	fmt.Println("\n2. Reading Nested Configuration:")
	fmt.Printf("   Server Host: %s\n", viper.GetString("server.host"))
	fmt.Printf("   Server Port: %d\n", viper.GetInt("server.port"))
	fmt.Printf("   Server Timeout: %d seconds\n", viper.GetInt("server.timeout"))

	// 3. Database configuration
	fmt.Println("\n3. Database Configuration:")
	fmt.Printf("   Driver: %s\n", viper.GetString("database.driver"))
	fmt.Printf("   Connection: %s:%d/%s\n",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"))

	fmt.Println("\n4. Cobra Commands Available:")
	fmt.Println("   Try running these commands:")
	fmt.Println("   go run cobra_main.go config show")
	fmt.Println("   go run cobra_main.go config validate")
	fmt.Println("   go run cobra_main.go --help")
	fmt.Println("   go run cobra_main.go config --help")
	fmt.Println("   go run cobra_main.go --app.debug config show")

	fmt.Println("\n5. Environment Variables (same as pflag version):")
	fmt.Println("   export VIPEREXAMPLE_APP_NAME=\"Cobra App\"")
	fmt.Println("   export VIPEREXAMPLE_SERVER_PORT=9999")
}

// showConfiguration displays current configuration
func showConfiguration() {
	fmt.Println("\n=== Current Configuration (via Cobra subcommand) ===")

	allKeys := viper.AllKeys()
	for _, key := range allKeys {
		fmt.Printf("%s: %v\n", key, viper.Get(key))
	}

	fmt.Printf("\nConfiguration loaded from: %s\n", viper.ConfigFileUsed())
}

// validateConfiguration validates current configuration
func validateConfiguration() {
	fmt.Println("\n=== Configuration Validation (via Cobra subcommand) ===")

	var errors []string

	// Validate required fields
	if viper.GetString("app.name") == "" {
		errors = append(errors, "app.name is required")
	}

	// Validate port range
	port := viper.GetInt("server.port")
	if port < 1 || port > 65535 {
		errors = append(errors, "server.port must be between 1 and 65535")
	}

	// Validate log level
	logLevel := viper.GetString("logging.level")
	validLevels := []string{"debug", "info", "warn", "error"}
	isValidLevel := false
	for _, level := range validLevels {
		if logLevel == level {
			isValidLevel = true
			break
		}
	}
	if !isValidLevel {
		errors = append(errors, fmt.Sprintf("logging.level must be one of: %v", validLevels))
	}

	if len(errors) == 0 {
		fmt.Println("✅ Configuration is valid")
		fmt.Printf("✅ All %d configuration keys are properly set\n", len(viper.AllKeys()))
	} else {
		fmt.Println("❌ Configuration validation failed:")
		for _, err := range errors {
			fmt.Printf("  - %s\n", err)
		}
	}
}
