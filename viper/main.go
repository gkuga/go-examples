package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// Initialize viper
	initConfig()

	// Demonstrate various ways to read configuration
	demonstrateViperFeatures()
}

// initConfig initializes viper configuration
func initConfig() {
	// Set default values
	setDefaults()

	// Setup command line flags
	setupFlags()

	// Configure config file settings
	viper.SetConfigName("config")                          // name of config file (without extension)
	viper.SetConfigType("yaml")                            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")                   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")                  // call multiple times to add many search paths
	viper.AddConfigPath(".")                               // optionally look for config in the working directory
	viper.SetEnvPrefix("VIPEREXAMPLE")                     // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // replace dots with underscores in env vars
	viper.AutomaticEnv()                                   // read in environment variables that match

	// Bind flags to viper
	viper.BindPFlags(pflag.CommandLine)

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Config file not found or error reading: %v\n", err)
	}
}

// setDefaults sets default configuration values
func setDefaults() {
	// App defaults
	viper.SetDefault("app.name", "Default Viper App")
	viper.SetDefault("app.version", "0.1.0")
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

// setupFlags defines command line flags
func setupFlags() {
	pflag.String("config", "", "config file (default is ./config.yaml)")
	pflag.String("app.name", "", "application name")
	pflag.String("server.host", "", "server host")
	pflag.Int("server.port", 0, "server port")
	pflag.Bool("app.debug", false, "enable debug mode")
	pflag.String("logging.level", "", "logging level (debug, info, warn, error)")
	pflag.Parse()
}

// demonstrateViperFeatures shows various viper capabilities
func demonstrateViperFeatures() {
	fmt.Println("\n=== Viper Configuration Demo ===")

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

	// 3. Reading with different data types
	fmt.Println("\n3. Database Configuration:")
	fmt.Printf("   Driver: %s\n", viper.GetString("database.driver"))
	fmt.Printf("   Connection: %s:%d/%s\n",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"))
	fmt.Printf("   User: %s\n", viper.GetString("database.user"))
	fmt.Printf("   SSL Mode: %s\n", viper.GetString("database.ssl_mode"))

	// 4. Reading entire sub-sections
	fmt.Println("\n4. Logging Configuration:")
	loggingConfig := viper.GetStringMapString("logging")
	for key, value := range loggingConfig {
		fmt.Printf("   %s: %s\n", key, value)
	}

	// 5. Reading arrays/slices (if available)
	fmt.Println("\n5. Features Configuration:")
	fmt.Printf("   Enable Metrics: %t\n", viper.GetBool("features.enable_metrics"))
	fmt.Printf("   Enable Tracing: %t\n", viper.GetBool("features.enable_tracing"))
	fmt.Printf("   Max Connections: %d\n", viper.GetInt("features.max_connections"))

	// 6. Demonstrate environment variable override
	fmt.Println("\n6. Environment Variable Examples:")
	fmt.Println("   Try setting environment variables like:")
	fmt.Println("   export VIPEREXAMPLE_APP_NAME=\"My Custom App\"")
	fmt.Println("   export VIPEREXAMPLE_SERVER_PORT=9999")
	fmt.Println("   export VIPEREXAMPLE_APP_DEBUG=true")

	// 7. Show all settings
	fmt.Println("\n7. All Configuration Keys:")
	allKeys := viper.AllKeys()
	for _, key := range allKeys {
		fmt.Printf("   %s: %v\n", key, viper.Get(key))
	}

	// 8. Demonstrate IsSet function
	fmt.Println("\n8. Checking if values are explicitly set:")
	checkKeys := []string{"app.name", "server.port", "nonexistent.key"}
	for _, key := range checkKeys {
		fmt.Printf("   %s is set: %t\n", key, viper.IsSet(key))
	}

	// 9. Demonstrate unmarshaling to struct
	fmt.Println("\n9. Unmarshaling to struct:")
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Error unmarshaling config: %v", err)
	} else {
		fmt.Printf("   Struct: %+v\n", config)
	}

	// 10. Writing configuration
	fmt.Println("\n10. Writing configuration to file:")
	viper.Set("runtime.last_run", "2024-01-01T12:00:00Z")
	viper.Set("runtime.run_count", viper.GetInt("runtime.run_count")+1)

	err = viper.WriteConfigAs("output_config.yaml")
	if err != nil {
		fmt.Printf("   Error writing config: %v\n", err)
	} else {
		fmt.Println("   Configuration written to output_config.yaml")
	}
}

// Config struct for unmarshaling
type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`

	Server struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"ssl_mode"`
	} `mapstructure:"database"`

	Logging struct {
		Level      string `mapstructure:"level"`
		File       string `mapstructure:"file"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAge     int    `mapstructure:"max_age"`
	} `mapstructure:"logging"`

	Features struct {
		EnableMetrics  bool `mapstructure:"enable_metrics"`
		EnableTracing  bool `mapstructure:"enable_tracing"`
		MaxConnections int  `mapstructure:"max_connections"`
	} `mapstructure:"features"`
}
