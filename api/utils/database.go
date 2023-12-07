package utils

import (
	"backend/api/config"
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// OpenDB abre la conexión con la base de datos y la devuelve
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBURL())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// loadEnv carga las variables de entorno
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
