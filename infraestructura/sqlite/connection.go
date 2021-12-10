package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewConn() (*sql.DB, error) {
	if _, err := os.Stat("./database.db"); os.IsNotExist(err) {
		log.Println("Creando base de datos database.db")
		file, err := os.Create("./database.db")
		if err != nil {
			return nil, err
		}
		file.Close()
		log.Println("database.db creado")
	}

	conn, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	if err := executeMigrations(conn); err != nil {
		return nil, err
	}

	return conn, err
}

func executeMigrations(db *sql.DB) error {
	beerTableSQL := `CREATE TABLE IF NOT EXISTS "beer" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name"	TEXT,
		"brewery"	TEXT,
		"country"	TEXT,
		"price"	NUMBER,
		"currency" TEXT);`
	stmt, err := db.Prepare(beerTableSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec()
	log.Println("Tablas: beer")
	return nil
}
