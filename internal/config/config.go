package config

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env   string     `yaml:"env"`
	GRPC  GRPCConfig `yaml:"grpc"`
	Minio MinioConfig
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type MinioConfig struct {
	Port              string
	MinioEndpoint     string
	MinioBucketName   string
	MinioRootUser     string
	MinioRootPassword string
	MinioUseSSL       bool
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("cfg path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("cfg file doesn't exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read cfg: " + err.Error())
	}
	cfg.Minio = *loadMinioConfig()

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to cfg file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func loadMinioConfig() *MinioConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env config: %v", err)
		return nil
	}

	return &MinioConfig{
		Port:              os.Getenv("PORT"),
		MinioEndpoint:     os.Getenv("MINIO_ENDPOINT"),
		MinioBucketName:   os.Getenv("MINIO_BUCKET_NAME"),
		MinioRootUser:     os.Getenv("MINIO_ROOT_USER"),
		MinioRootPassword: os.Getenv("MINIO_ROOT_PASSWORD"),
		MinioUseSSL:       getEnvAsBool("MINIO_USE_SSL"),
	}
}

func getEnvAsBool(key string) bool {
	if strVal := os.Getenv(key); strVal != "" {
		if val, err := strconv.ParseBool(strVal); err == nil {
			return val
		}
	}
	return false // возвращаю по дефолту false если значение не задано
}
