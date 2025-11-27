package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

/*
	Base no "https://go.dev/doc/tutorial/database-access"
	Queria aprender a usar o MySQL um Banco de dados
	simples que muita gente usa e tem uma maturidade
	bem grande
*/

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "meu_SQL"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	println("Deu certo!\nebaaaaaaaaaaaa!!!")
}
