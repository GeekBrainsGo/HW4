/*
Go. Homework 3
Zaur Malakhov, dated Sep 21, 2019
*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/MySQL"
)

func main() {
	r := chi.NewRouter()
	lg := logrus.New()

	db, err := sql.Open("mysql", "mysql:root@/blog")
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}
	defer db.Close()

	serv := Server{
		lg:    lg,
		db:    db,
		Title: "БЛОГ О СПОРТЕ",
		Posts: Posts{},
	}

	r.Route("/", func(r chi.Router) {
		r.Get("/", serv.HandleGetIndexHtml)
		r.Get("/view/{postID}", serv.HandleGetPostHtml)
		r.Get("/edit/{postID}", serv.HandleGetEditHtml)
		r.Get("/new", serv.HandleGetNewHtml)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/new", serv.postNewHandler)
			r.Delete("/delete/{id}", serv.deletePostHandler)
			r.Put("/update/{id}", serv.putPostHandler)
		})
	})

	lg.Info("server is starts")
	http.ListenAndServe(":8080", r)
}

type Server struct {
	lg    *logrus.Logger
	db    *sql.DB
	Title string
	Posts Posts
}

type Posts []Post
type Post struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Date             string `json:"date"`
	SmallDescription string `json:"smalldescription"`
	Description      string `json:"description"`
}

// Insert - добавляет пост в БД
func (post *Post) Insert(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO posts (ID, title, date, smalldescription, description) VALUES (?, ?, ?, ?, ?)",
		post.ID, post.Title, post.Date, post.SmallDescription, post.Description,
	)
	return err
}

// Delete - удалят объект из базы
func (post *Post) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM posts WHERE ID = ?",
		post.ID,
	)
	return err
}

// Update - изменяет пост в БД
func (post *Post) Update(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE posts SET title = ?, date = ?, smalldescription = ?, description = ? WHERE ID = ?",
		post.Title, post.Date, post.SmallDescription, post.Description, post.ID,
	)
	return err
}

func (serv *Server) HandleGetIndexHtml(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/index.html")
	data, _ := ioutil.ReadAll(file)

	posts, errAllPosts := GetAllPosts(serv.db)
	if errAllPosts != nil {
		serv.lg.WithError(errAllPosts).Error("GetAllPosts")
	}
	serv.Posts = posts

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", serv)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) HandleGetPostHtml(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")

	postFindId := -1
	for i := 0; i < len(serv.Posts); i++ {
		if serv.Posts[i].ID == postIDStr {
			postFindId = i
		}
	}
	post := serv.Posts[postFindId]

	file, _ := os.Open("./www/static/post.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) HandleGetEditHtml(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")

	postFindId := 0
	for i := 0; i < len(serv.Posts); i++ {
		if serv.Posts[i].ID == postIDStr {
			postFindId = i
		}
	}
	post := serv.Posts[postFindId]

	file, _ := os.Open("./www/static/edit.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", post)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) HandleGetNewHtml(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/new.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", serv)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

// postNewHandler - добавление нового поста
func (serv *Server) postNewHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	post := Post{}
	_ = json.Unmarshal(data, &post)

	post.ID = uuid.NewV4().String()

	if err := post.Insert(serv.db); err != nil {
		//serv.SendInternalErr(w, err)
		serv.lg.WithError(err).Error("DB insert")
		return
	}

	data, _ = json.Marshal(post)
	w.Write(data)
}

// deletePostHandler - удаляем пост
func (serv *Server) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	post := Post{ID: postID}

	if err := post.Delete(serv.db); err != nil {
		serv.lg.WithError(err).Error("DB delete")
		return
	}
}

// putPostHandler - обновляем пост
func (serv *Server) putPostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	fmt.Println(postID)

	data, _ := ioutil.ReadAll(r.Body)

	post := Post{}
	_ = json.Unmarshal(data, &post)
	post.ID = postID

	data, _ = json.Marshal(post)
	w.Write(data)

	if err := post.Update(serv.db); err != nil {
		serv.lg.WithError(err).Error("DB update")
		return
	}

}

// GetAllPosts - получение всех задач
func GetAllPosts(db *sql.DB) (Posts, error) {
	rows, err := db.Query("SELECT ID, title, date, smalldescription, description  FROM posts")
	if err != nil {
		return nil, err
	}
	posts := make(Posts, 0, 8)
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.Title, &post.Date, &post.SmallDescription, &post.Description); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}
