package main

/*
	Basics Go.
	Rishat Ishbulatov, dated Sep 25, 2019.
	Create models for your structures in the database.
	Create methods to get data from the database on your models.
	Adapt routes that handle requests for all posts, a specific 
	blog post, and edit pages.
*/

import (
	"HW4/sqlblog/server"
	"database/sql"
	"flag"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/MySQL"
	"github.com/sirupsen/logrus"
)

func main() {
	flagRootDir := flag.String("rootdir", "./www", "root dir of the server")
	flagServAddr := flag.String("addr", "localhost:8080", "server address")
	flag.Parse()

	lg := NewLogger()
	db, err := sql.Open("mysql", "mysql:root@/sqlblog")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}
	defer db.Close()
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
}

// NewLogger creates new logger.
func NewLogger() *logrus.Logger {
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	lg.SetLevel(logrus.DebugLevel)
	return lg
}
