package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"serv/models"

	"github.com/go-chi/chi"
)

func (serv *Server) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/index.html")
	data, _ := ioutil.ReadAll(file)
	if blogItems, err := models.GetAllBlogItems(serv.db); err != nil {
		serv.lg.Error("Error getting all posts", err)
	} else {
		indexTemplate := template.Must(template.New("index").Parse(string(data)))
		err := indexTemplate.ExecuteTemplate(w, "index", blogItems)
		if err != nil {
			serv.lg.WithError(err).Error("template")
		}
	}

}

func (serv *Server) handleGetPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/post.html")
	data, _ := ioutil.ReadAll(file)
	postNumberStr := chi.URLParam(r, "id")
	indexTemplate := template.Must(template.New("index").Parse(string(data)))
	var searchedPost models.BlogItem
	searchedPost, err := models.GetSingleBlogItem(postNumberStr, serv.db)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	err = indexTemplate.ExecuteTemplate(w, "index", searchedPost)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) handleGetEditPost(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/edit.html")
	data, _ := ioutil.ReadAll(file)
	postNumberStr := chi.URLParam(r, "id")
	indexTemplate := template.Must(template.New("index").Parse(string(data)))
	var searchedPost models.BlogItem
	searchedPost, err := models.GetSingleBlogItem(postNumberStr, serv.db)
	if err != nil {
		serv.lg.Error("Error getting post", err)
	}
	err = indexTemplate.ExecuteTemplate(w, "index", searchedPost)
	if err != nil {
		serv.lg.WithError(err).Error("template")
	}
}

func (serv *Server) handlePostEditPost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogItem
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Update(serv.db)
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}

}

func (serv *Server) handlePostCreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogItem
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Update(serv.db)
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}
}

func (serv *Server) handlePostDeletePost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogItem
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		post.Delete(serv.db)
		resp, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}

}
