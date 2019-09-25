package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"database/sql"
	_ "github.com/go-sql-driver/MySQL"
)

const (
	STATIC_DIR = "./www/"
)

func main() {
	r := chi.NewRouter()
	lg := logrus.New()
	db, err := sql.Open("mysql", "homestead:secret@tcp(127.0.0.1:33060)/blog")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}

	serv := Server{
		lg:    lg,
		db:    db,
		Title: "BLOG",
		Posts: PostItemSlice {
			{
				Id:		   1,
				Title:     "Пост 1",
				Text:      "Очень интересный текст",
				Labels:    []string{"путешестве", "отдых"},
			},
			{
				Id:		   2,
				Title:     "Пост 2",
				Text:      "Второй очень интересный текст",
				Labels:    []string{"домашка", "golang"},
			},
			{
				Id:		   3,
				Title:     "Пост 3",
				Text:      "Третий очень интересный текст",
				Labels:    []string{},
			},
		},
	}

	fileServer := http.FileServer(http.Dir(STATIC_DIR))
	r.Handle("/static/*", fileServer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", serv.HandleGetIndex)
		r.Get("/post/{id}", serv.HandleGetPost)
		r.Get("/post/create", serv.HandleGetEditPost)
		r.Get("/post/{id}/edit", serv.HandleGetEditPost)
	})

	r.Route("/api/v1/", func(r chi.Router) {
		r.Post("/post/create", serv.HandleEditPost)
		r.Post("/post/{id}/edit", serv.HandleEditPost)
	})

	lg.Info("server is start")
	http.ListenAndServe(":8080", r)
}

type Server struct {
	lg    *logrus.Logger
	db    *sql.DB
	Title string
	Posts PostItemSlice
}

func (serv *Server) AddOrUpdatePost(newPost PostItem) (PostItem) {
	for key, post := range serv.Posts {
		if post.Id == newPost.Id {
			serv.Posts[key] = newPost
			return post
		}
	}

	err := newPost.Update(serv.db)
	if err != nil {
		newPost, err = newPost.Create(serv.db)
	}

	return newPost
}


func (serv *Server) HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/templates/index.gohtml")
	data, _ := ioutil.ReadAll(file)

	posts, err := GetAllPostItems(serv.db)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetIndex")
	}

	serv.Posts = posts

	templ := template.Must(template.New("page").Parse(string(data)))
	err = templ.ExecuteTemplate(w, "page", serv)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetIndexTemplate")
	}
}

func (serv *Server) HandleGetPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/templates/post.gohtml")
	data, _ := ioutil.ReadAll(file)

	postIDStr := chi.URLParam(r, "id")
	postID, _ := strconv.ParseInt(postIDStr, 10, 64)

	post, err := GetPostItem(postID, serv.db)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetPost")
		post = PostItem{}
	}

	templ := template.Must(template.New("page").Parse(string(data)))
	err = templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetPostTemplate")
	}
}

func (serv *Server) HandleGetEditPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/templates/edit_post.gohtml")
	data, _ := ioutil.ReadAll(file)

	postIDStr := chi.URLParam(r, "id")
	postID, _ := strconv.ParseInt(postIDStr, 10, 64)

	post, err := GetPostItem(postID, serv.db)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetEditPost")
		post = PostItem{}
	}

	templ := template.Must(template.New("page").Parse(string(data)))
	err = templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("HandleGetEditPostTemplate")
	}
}

func (serv *Server) HandleEditPost(w http.ResponseWriter, r *http.Request) {

	/*
	{"id":4, "Title":"Пост 4", "Text":"Новый очень интересный текст", "Labels":["l1","l2"]}
	*/

	decoder := json.NewDecoder(r.Body)
	var inPostItem PostItem
	err := decoder.Decode(&inPostItem)
	if err != nil {
		serv.lg.WithError(err).Error("HandleEditPost")
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	newPost := serv.AddOrUpdatePost(inPostItem)
	respondWithJSON(w, http.StatusOK, newPost)
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

