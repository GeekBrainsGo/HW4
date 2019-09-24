/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 22.09.2019 13:11
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func main() {

	// connect to DB
	db, err := sql.Open("mysql", myCnf("client")+DSN)
	if err != nil {
		log.Fatalln("Can't open DB:", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	// 1st connect wil be here
	if err = db.Ping(); err != nil {
		log.Fatalln("Can't ping DB:", err)
	}

	// new handlers struct
	handlers := &Handler{
		tmplGlob: template.Must(template.ParseGlob(path.Join(TEMPLATEPATH, TEMPLATEEXT))),
		db:       db,
	}

	// prepare server, routes & middleware
	srv := &http.Server{Addr: SERVADDR, Handler: handlers.prepareRoutes()}

	// graceful shutdown
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) // os.Kill cannot be trapped anyway!
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Signal received:", <-shutdown)
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Error while shutdown server:", err)
		}
	}()

	fmt.Println("Starting server at:", SERVADDR)
	log.Printf("Shutdown server at: %s\n%v", SERVADDR, srv.ListenAndServe())
}

// read MySQL parameters from .my.cnf
func myCnf(profile string) string {
	cnf := path.Join(os.Getenv("HOME"), ".my.cnf")
	cfg, err := ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true}, cnf)
	if err != nil {
		return ""
	}
	for _, s := range cfg.Sections() {
		if s.Name() != profile {
			continue
		}
		user := s.Key("user")
		password := s.Key("password")
		host := s.Key("host")
		port := s.Key("port")
		return fmt.Sprintf("%s:%s@tcp(%s:%s)", user, password, host, port)
	}
	return ""
}
