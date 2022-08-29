package utils

import (
	"os"
)

type EnvironmentVariables struct {
	PORT                 string
	DB_CONNECTION_STRING string
	AUTH_JWKS_URI        string
}

var Env *EnvironmentVariables = &EnvironmentVariables{}

func ConfigureEnv() {
	// var err error
	if Env.PORT = os.Getenv("PORT"); Env.PORT == "" {
		Env.PORT = "8080"
	}
	if Env.DB_CONNECTION_STRING = os.Getenv("DB_CONNECTION_STRING"); Env.DB_CONNECTION_STRING == "" {
		Env.DB_CONNECTION_STRING = "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
	}
	if Env.AUTH_JWKS_URI = os.Getenv("AUTH_JWKS_URI"); Env.AUTH_JWKS_URI == "" {
		Env.AUTH_JWKS_URI = "https://somedomain.com"
	}
}
