package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	SupabaseURL string
	SupabaseKey string
	Port        string
	BaseUrl     string
	Domain      string
}

func NewConfig() Config {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("Error loading dotenv file: %s\n", err)
			os.Exit(1)
		}
	}
	for _, env := range []string{"SUPABASE_URL", "SUPABASE_KEY", "PORT", "BASE_URL", "DOMAIN"} {
		if os.Getenv(env) == "" {
			fmt.Printf("Environment variable %s not found\n", env)
			os.Exit(1)
		}
	}
	return Config{
		Env:         os.Getenv("ENV"),
		SupabaseURL: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_KEY"),
		Port:        os.Getenv("PORT"),
		BaseUrl:     os.Getenv("BASE_URL"),
		Domain:      os.Getenv("DOMAIN"),
	}
}

func (c Config) WithSupabaseURL(url string) Config {
	c.SupabaseURL = url
	return c
}

func (c Config) WithSupabaseKey(key string) Config {
	c.SupabaseKey = key
	return c
}

func (c Config) WithPort(port string) Config {
	c.Port = port
	return c
}
