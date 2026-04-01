package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		connStr = "postgresql://admin:q24RE6widrfPyqL3en21N1HLuSwkXUmm@dpg-d76hlq7pm1nc73967v1g-a.oregon-postgres.render.com/mayorgasgarage"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la conexión: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")
}
