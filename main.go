package main

import (
	"database/sql"
	"flag"
	"myblog/server"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/MySQL"
)

func main() {
	flagRootDir := flag.String("rootdir", "./web", "root dir of the server")
	flagServAddr := flag.String("addr", "localhost:8080", "server address")
	flag.Parse()

	lg := NewLogger()
	// db, err := sql.Open("mysql", "mysql:root@/MyBlogs")
	db, err := sql.Open("mysql", "root:root@/MyBlogs")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	serv := server.New(lg, *flagRootDir, db)

	go func() {
		err := serv.Start(*flagServAddr)
		if err != nil {
			lg.WithError(err).Fatal("can't run the server")
		}
	}()

	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig
	lg.Info("Stop server!")

}
