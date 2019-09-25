package server

import (
	"HW4/sqlblog/models"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/go-chi/chi"
)

// templateHandle returns template.
func (s *Server) templateHandle(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = s.indexTemplate
	}

	file, err := os.Open(path.Join(s.rootDir, s.templatesDir, templateName))
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.SendInternalErr(w, err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	templ, err := template.New("").Parse(string(data))
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	posts, err := models.AllPosts(s.db)
	if err != nil {
		s.SendInternalErr(w, err)
		return
	}

	s.Page.Posts = posts

	if err := templ.Execute(w, s.Page); err != nil {
		s.SendInternalErr(w, err)
		return
	}
}

// postHandle adds new post.
func (s *Server) postHandle(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	post := models.Post{}
	_ = json.Unmarshal(data, &post)

	if err := post.Insert(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(post)
	w.Write(data)
}

// deleteHandle deletes a post.
func (s *Server) deleteHandle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post := models.Post{ID: id}
	if err := post.Delete(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}
}

// putHandle renew post.
func (s *Server) putHandle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, _ := ioutil.ReadAll(r.Body)

	post := models.Post{}
	_ = json.Unmarshal(data, &post)
	post.ID = id

	if err := post.Update(s.db); err != nil {
		s.SendInternalErr(w, err)
		return
	}
}
