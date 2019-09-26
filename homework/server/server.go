package server

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

//Server - main server structure
type Server struct {
	lg      *logrus.Logger
	db      *sql.DB
	rootDir string
}

//New - creates and returns new server
func New(lg *logrus.Logger, db *sql.DB) *Server {
	return &Server{
		lg: lg,
		db: db,
	}
}

//Start - starts the server
func (serv *Server) Start(addr string) error {
	r := chi.NewRouter()
	serv.bindRoutes(r)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "www/static")
	FileServer(r, "/static", http.Dir(filesDir))

	serv.lg.Debug("Server was started")
	return http.ListenAndServe(addr, r)
}

//FileServer starts file server for static contents
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
