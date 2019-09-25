package webserver

import (
	"blog/app/database"
	"blog/app/database/mysql"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/MySQL"
	"github.com/sirupsen/logrus"
)

// WebServer ...
type WebServer struct {
	router   *chi.Mux
	logger   *logrus.Logger
	database database.Database
}

func newServer(db database.Database) *WebServer {
	serv := &WebServer{
		router:   chi.NewRouter(),
		logger:   logrus.New(),
		database: db,
	}

	serv.configureRouter()

	return serv
}

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseConnectionString)
	if err != nil {
		return err
	}

	defer db.Close()
	sqlDatabase := mysql.New(db)
	serv := newServer(sqlDatabase)
	return http.ListenAndServe(config.BindAddr, serv)
}

func newDB(dsnURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsnURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (serv *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv *WebServer) configureRouter() {
	//routes
	serv.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	serv.router.HandleFunc("/list", serv.postListHandle())

	serv.router.HandleFunc("/view/{postID}", serv.postViewHandle())

}

func (serv *WebServer) postListHandle() http.HandlerFunc {

	type PageModel struct {
		Title string
		Data  interface{}
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := serv.database.Post().FindAll()
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		pageData := PageModel{
			Title: "Posts List",
			Data:  data,
		}

		templ := template.Must(template.New("page").ParseFiles("./templates/blog/List.tpl", "./templates/common.tpl"))
		err = templ.ExecuteTemplate(w, "page", pageData)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

	}
}

func (serv *WebServer) postViewHandle() http.HandlerFunc {

	type PageModel struct {
		Title string
		Data  interface{}
	}

	return func(w http.ResponseWriter, r *http.Request) {

		postIDStr := chi.URLParam(r, "postID")
		postID, _ := strconv.ParseInt(postIDStr, 10, 64)

		data, err := serv.database.Post().Find(postID)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

		pageData := PageModel{
			Title: "View Post",
			Data:  data,
		}

		templ := template.Must(template.New("page").ParseFiles("./templates/blog/View.tpl", "./templates/common.tpl"))
		err = templ.ExecuteTemplate(w, "page", pageData)
		if err != nil {
			serv.errorAPI(w, r, http.StatusInternalServerError, err)
			return
		}

	}
}

func (serv *WebServer) errorAPI(w http.ResponseWriter, r *http.Request, code int, err error) {
	serv.respondJSON(w, r, code, map[string]string{"error": err.Error()})
}

func (serv *WebServer) respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (serv *WebServer) respondWhithTemplate(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
