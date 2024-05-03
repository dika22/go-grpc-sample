package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env              string
	App              App
	// Database         database.Config
	// DatabaseRpt      database.Config
	// Kafka            kafka.Config
	// Log              logger.LogConfig
	GrpcServer       Grpc
}

type App struct {
	Name     string
	Host     string
	Version  string
	Timezone string
	Port     int
}

type Grpc struct {
	Name string
	Port string
	Host string
}

func LoadENVConfig() (EnvironmentConfig, error) {
	err := godotenv.Load(dir(".env"))
	if err != nil {
		err = fmt.Errorf("failed to load .env file: %w", err)
		return EnvironmentConfig{}, err
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf("error when convert string to int: %w", err)
		return EnvironmentConfig{}, err
	}

	config := EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:     os.Getenv("APP_NAME"),
			Host:     os.Getenv("APP_HOST"),
			Version:  os.Getenv("APP_VERSION"),
			Timezone: os.Getenv("APP_TIMEZONE"),
			Port:     port,
		},
		// Database: database.Config{
		// 	Dialect:  os.Getenv("DB_DIALECT"),
		// 	Host:     os.Getenv("DB_HOST"),
		// 	Port:     os.Getenv("DB_PORT"),
		// 	Name:     os.Getenv("DB_NAME"),
		// 	Username: os.Getenv("DB_USERNAME"),
		// 	Password: os.Getenv("DB_PASSWORD"),
		// 	Timezone: os.Getenv("DB_TIMEZONE"),
		// },
		
		GrpcServer: Grpc{
			Name: os.Getenv("GRPC_SERVER_NAME"),
			Port: os.Getenv("GRPC_SERVER_PORT"),
			Host: os.Getenv("GRPC_SERVER_HOST"),
		},
		// GrpcClientPu: Grpc{
		// 	Name: os.Getenv("GRPC_PU_NAME"),
		// 	Port: os.Getenv("GRPC_PU_PORT"),
		// 	Host: os.Getenv("GRPC_PU_HOST"),
		// },
	}

	return config, nil
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
