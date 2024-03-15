package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgreSQLDB() (*sql.DB, error) {

	host := "34.41.188.11"
	port := "5432"
	dbname := "starDB"
	user := "admin"
	password := "123456789!!!!!!!"

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error de conexión a la base de datos:", err)
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Satellites (
    name TEXT PRIMARY KEY NOT NULL,
    distance NUMERIC,
    message TEXT
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla Satellites:", err)
		db.Close()
		return nil, err
	}

	return db, nil
}


