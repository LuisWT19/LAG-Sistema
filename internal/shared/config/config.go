package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config alamacena toda la configuracion de la aplicación.
type Config struct {
	AppName string
	AppEnv  string
	AppPort string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	JWTSecret  string
	JWTExpires string
}

var cfg *Config

// Load carga las variables de entorno desde el archivo .env
func Load() {
	//Carga el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("No se pudo cargar el archivo .env")
	}

	cfg = &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTExpires: os.Getenv("JWT_EXPIRES"),
	}
}

// Get devuelve la configuración cargada
func Get() *Config {
	return cfg
}
