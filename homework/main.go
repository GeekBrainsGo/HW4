package main

import (
	"database/sql"
	"flag"
	"serv/server"

	_ "github.com/go-sql-driver/MySQL"
)

func main() {
	flagServAddr := flag.String("addr", "localhost:8080", "server address")

	lg := NewLogger()
	db, err := sql.Open("mysql", "mysql:mysql123!@tcp(192.168.99.100:3306)/blog")
	if err != nil {
		lg.Panic("Can't connect to DB", err)
	} else {
		lg.Info("Connection to DB successful")
	}
	lg.Info(db.Ping())
	srv := server.New(lg, db)
	srv.Start(*flagServAddr)
}
