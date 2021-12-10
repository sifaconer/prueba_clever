package main

import (
	"log"

	server "prueba.clever.com/infraestructura/http"
	"prueba.clever.com/infraestructura/sqlite"
)

const PORT = "8182"

func main() {

	db, err := sqlite.NewConn()
	if err != nil {
		log.Fatal("Error to connect database", err)
	}

	s := server.Server{
		Port: PORT,
		Repository: &sqlite.BeersImpl{
			DB: db,
		},
	}

	s.New()
}
