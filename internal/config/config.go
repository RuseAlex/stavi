package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"stavi/internal/logger"
	"strconv"
)

type Config struct {
	Debug bool
	Port  int
	api   string
	db    struct {
		dsn string
	}
	stripe struct {
		key    string
		secret string
	}
	email struct {
		sender   string
		password string
		host     string
		port     string
	}
}

type ConfigErr struct {
	Err   error
	Level logger.LogLevel
}

func (cfg *Config) LoadEnv(envFile string) ConfigErr {
	// try to read the .env file in settings
	var env map[string]string
	var err error

	// if envFile location is not an empty string then read it, else resort to the default location
	if envFile != "" {
		env, err = godotenv.Read(envFile)
		if err != nil {
			return ConfigErr{Err: err, Level: logger.FATAL}
		}
	}

	env, err = godotenv.Read("./settings/.env")
	if err != nil {
		return ConfigErr{Err: err, Level: logger.FATAL}
	}

	// try to load app details into the config from env
	cfg.Port, err = strconv.Atoi(env["PORT"])
	if err != nil {
		return ConfigErr{Err: err, Level: logger.FATAL}
	}

	// try to load database details into the config from env
	user := env["DB_USER"]
	password := env["DB_PASSWORD"]
	dbname := env["DB_NAME"]
	host := env["DB_HOST"]
	sslmode := env["SSL_MODE"]
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
	cfg.stripe.key = env["STRIPE_KEY"]
	cfg.stripe.secret = env["STRIPE_SECRET"]
	if cfg.stripe.key == "" || cfg.stripe.secret == "" {
		return ConfigErr{Err: errors.New("missing STRIPE env variables"), Level: logger.WARNING}
	}

	// try to load email details into the config from env
	cfg.email.port = env["EMAIL_PORT"]
	cfg.email.host = env["EMAIL_HOST"]
	cfg.email.sender = env["EMAIL_HOST"]
	cfg.email.password = env["EMAIL_PASSWORD"]
	if cfg.email.host == "" || cfg.email.sender == "" || cfg.email.password == "" || cfg.email.port == "" {
		return ConfigErr{Err: errors.New("missing EMAIL env variables"), Level: logger.WARNING}
	}

	return ConfigErr{Err: nil, Level: logger.DEBUG}
}
