package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	DbConfig struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
		// Additional configurations⤵️
	}

	ApiConfig struct {
		ApiHost string
		ApiPort string
		// Additional configurations⤵️
	}

	FileLogConfig struct {
		FilePath string
		// Additional configurations⤵️
	}

	JWTConfig struct {
		Issuer           string
		SignatureKey     []byte
		SigningMethod    *jwt.SigningMethodHMAC
		ExpiresInMinutes int
		// Additional configurations⤵️
	}
)

type Config struct {
	DbConfig
	ApiConfig
	FileLogConfig
	JWTConfig
}

func (c *Config) LoadConfig() error {
	v := viper.New()
	v.AutomaticEnv()

	c.DbConfig = DbConfig{
		Driver:   v.GetString("APP_DB_DRIVER"),
		Host:     v.GetString("APP_DB_HOST"),
		Port:     v.GetString("APP_DB_PORT"),
		User:     v.GetString("APP_DB_USER"),
		Password: v.GetString("APP_DB_PASSWORD"),
		Name:     v.GetString("APP_DB_NAME"),
		SSLMode:  v.GetString("APP_SSL_MODE"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: v.GetString("APP_API_HOST"),
		ApiPort: v.GetString("APP_API_PORT"),
	}

	c.FileLogConfig = FileLogConfig{
		FilePath: v.GetString("APP_FILE_PATH"),
	}

	c.JWTConfig = JWTConfig{
		Issuer:           v.GetString("APP_ISSUER"),
		SignatureKey:     []byte(v.GetString("APP_SIGNATURE_KEY")),
		SigningMethod:    jwt.SigningMethodHS512,
		ExpiresInMinutes: v.GetInt("APP_EXPIRES_IN_MINUTES"),
	}

	if c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Name == "" || c.DbConfig.SSLMode == "" ||
		c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" ||
		c.FileLogConfig.FilePath == "" ||
		c.JWTConfig.Issuer == "" || len(c.JWTConfig.SignatureKey) == 0 || c.JWTConfig.ExpiresInMinutes == 0 {
		return fmt.Errorf("missing required env var")
	}
	return nil
}

// NewConfig Constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
