package server

import (
	"database/sql"
	"encoding/json"
	"myblog/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Server - объект сервера
type Server struct {
	lg            *logrus.Logger
	db            *sql.DB
	rootDir       string
	templatesDir  string
	indexTemplate string
	Page          models.Page
}

// New - создаёт новый экземпляр сервера
func New(lg *logrus.Logger, rootDir string, db *sql.DB) *Server {
	return &Server{
		lg:            lg,
		db:            db,
		rootDir:       rootDir,
		templatesDir:  "/templates",
		indexTemplate: "index.html",
		Page: models.Page{
			Blogs: models.BlogItemSlice{},
			// Page: models.Page{
			// 	Tasks: models.TaskItemSlice{
			// 		{ID: "0", Text: "123", Completed: false},
			// 		{ID: "1", Text: "test", Completed: true},
			// 		{ID: "2", Text: "test 2", Completed: false},
			// 	},
			// },
		},
	}
}

// Start - запускает сервер
func (serv *Server) Start(addr string) error {
	r := chi.NewRouter()
	serv.bindRoutes(r)
	serv.lg.Debug("server is started ...")
	return http.ListenAndServe(addr, r)
}

func (serv *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", serv.getTemplateHandler)
		r.Route("/blog", func(r chi.Router) {
			r.Get("/{id}", serv.viewBlogHandler)
			r.Get("/del/{id}", serv.deleteBlogHandler)
			// 	r.Post("/tasks", serv.postTaskHandler)
			// 	r.Delete("/tasks/{id}", serv.deleteTaskHandler)
			// 	r.Put("/tasks/{id}", serv.putTaskHandler)
		})
	})
}

// SendErr - отправляет ошибку пользователю и логирует её
func (serv *Server) SendErr(w http.ResponseWriter, err error, code int, obj ...interface{}) {
	serv.lg.WithField("data", obj).WithError(err).Error("server error")
	w.WriteHeader(code)
	errModel := models.ErrorModel{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}
	data, _ := json.Marshal(errModel)
	w.Write(data)
}

// SendInternalErr - отправляет 500 ошибку
func (serv *Server) SendInternalErr(w http.ResponseWriter, err error, obj ...interface{}) {
	serv.SendErr(w, err, http.StatusInternalServerError, obj)
}
