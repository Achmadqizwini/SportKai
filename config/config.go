package config

import (
	"os"
	"strconv"
)

func GetConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			AppHost: getEnv("APP_HOST", "localhost"),
			AppPort: getEnvAsInt("APP_PORT", 8000),
		},
		DBconfig: DBconfig{
			DB_DRIVER:   getEnv("DB_DRIVER", "mysql"),
			DB_HOST:     getEnv("DB_HOST", "localhost"),
			DB_PORT:     getEnvAsInt("DB_PORT", 3306),
			DB_USERNAME: getEnv("DB_USER", "root"),
			DB_PASSWORD: getEnv("DB_PASSWORD", ""),
			DB_NAME:     getEnv("DB_NAME", "go-simple-template"),
		},
		CacheConfig: CacheConfig{
			CacheDriver: getEnv("CACHE_DRIVER", "redis"),
			Redis: Redis{
				RedisHost:     getEnv("REDIS_HOST", "localhost"),
				RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
				RedisDB:       getEnvAsInt("REDIS_DB", 0),
				RedisPassword: getEnv("REDIS_PASSWORD", ""),
			},
		},
		StorageConfig: StorageConfig{
			StorageDriver: getEnv("STORAGE_DRIVER", "minio"),
			Minio: Minio{
				MinioEndpoint:        getEnv("MINIO_ENDPOINT", ""),
				MinioAccessKeyID:     getEnv("MINIO_ACCESS_KEY_ID", ""),
				MinioAccessKeySecret: getEnv("MINIO_ACCESS_KEY_SECRET", ""),
				MinioBucketName:      getEnv("MINIO_BUCKET_NAME", ""),
			},
			GCS: GCS{
				CredentialsFile: getEnv("GCS_CREDENTIALS_FILE", ""),
				GCSBucketName:   getEnv("GCS_BUCKET_NAME", ""),
			},
		},
	}
}

type Config struct {
	AppConfig
	DBconfig
	CacheConfig
	StorageConfig
}

type AppConfig struct {
	AppName    string
	AppVersion string
	AppHost    string
	AppPort    int
}

type DBconfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
}

type CacheConfig struct {
	CacheDriver string
	Redis       Redis
}

type StorageConfig struct {
	StorageDriver string
	Minio         Minio
	GCS           GCS
}

type Minio struct {
	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioAccessKeySecret string
	MinioBucketName      string
}

type GCS struct {
	CredentialsFile string
	GCSBucketName   string
}

type Redis struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
