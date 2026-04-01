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
		log.Println("⚠️ DATABASE_URL no definida")
		return
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("❌ Error al abrir conexión:", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Println("❌ Error al hacer ping a la DB:", err)
		return
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL")
}
