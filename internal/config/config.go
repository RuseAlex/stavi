package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"stavi/internal/logger"
)

type Config struct {
	Debug bool
	port  int
	api   string
	db    struct {
		dsn string
	}
	stripe struct {
		key    string
		secret string
	}
	env map[string]string
}

type ConfigErr struct {
	Err   error
	Level logger.LogLevel
}

func (cfg *Config) LoadEnv() ConfigErr {
	// try to read the .env file in settings
	var err error
	cfg.env, err = godotenv.Read("./settings/.env")
	if err != nil {
		return ConfigErr{Err: err, Level: logger.FATAL}
	}

	// try to load database details into the config from env
	user := cfg.env["DB_USER"]
	password := cfg.env["DB_PASSWORD"]
	dbname := cfg.env["DB_NAME"]
	host := cfg.env["DB_HOST"]
	sslmode := cfg.env["SSL_MODE"]
	// check to see if any of these are empty
	if user == "" || password == "" || dbname == "" || host == "" || sslmode == "" {
		return ConfigErr{Err: errors.New("missing DB env variables"), Level: logger.WARNING}
	}

	cfg.db.dsn = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s",
		user,
		password,
		dbname,
		host,
		sslmode)

	// try to load stripe details into the config from env
	cfg.stripe.key = cfg.env["STRIPE_KEY"]
	cfg.stripe.secret = cfg.env["SECRET"]

	// try to load email details in the config from env

	return ConfigErr{}
}

func (cfg *Config) Add(envName, envValue string) error {
	if _, exists := cfg.env[envName]; !exists {
		cfg.env[envName] = envValue
		return nil
	}

	return errors.New("error: can't add new env, an env with the same name exist already")
}

func (cfg *Config) Get(envName string) string {
	envValue, ok := cfg.env[envName]
	if !ok {
		return ""
	}

	return envValue
}
