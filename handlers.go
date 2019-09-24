/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 22.09.2019 13:11
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"log"
	"net/http"
)

// Handler is the global server handlers struct.
type Handler struct {
	db       *sql.DB
	posts    dbPosts
	tmplGlob *template.Template
	Error
}

// main page
func (h *Handler) mainPageForm(w http.ResponseWriter, _ *http.Request) {
	var err error
	h.posts, err = h.posts.getPosts("", h.db)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := h.tmplGlob.ExecuteTemplate(w, "index", h.posts); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// one post page
func (h *Handler) postsPageForm(w http.ResponseWriter, r *http.Request) {
	postNum := r.URL.Query().Get("id")
	if postNum == "" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	posts, err := h.posts.getPosts(postNum, h.db)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	posts[0].Body = template.HTML(blackfriday.Run([]byte(posts[0].Body)))
	h.execTemplate(w, posts[0], "post")
}

// create post page
func (h *Handler) createPostPageForm(w http.ResponseWriter, _ *http.Request) {
	h.execTemplate(w, Post{}, "create")
}

// edit post page
func (h *Handler) editPostPageForm(w http.ResponseWriter, r *http.Request) {
	posts, err := h.posts.getPosts(r.URL.Query().Get("id"), h.db)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.execTemplate(w, posts[0], "edit")
}

// exec template helper
func (h *Handler) execTemplate(w http.ResponseWriter, post Post, tmpl string) {
	if err := h.tmplGlob.ExecuteTemplate(w, tmpl, post); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// api create post
func (h *Handler) createPostPage(w http.ResponseWriter, r *http.Request) {
	post := h.decodePost(w, r)
	if post == nil {
		return
	}
	if err := h.posts.createPost(post, h.db); err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "error while create post")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// api edit post
func (h *Handler) editPostPage(w http.ResponseWriter, r *http.Request) {
	postNum := chi.URLParam(r, "id")
	post := h.decodePost(w, r)
	if post == nil {
		return
	}
	post.ID = postNum
	err := h.posts.updatePost(post, h.db)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "error while update post")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// api delete post
func (h *Handler) deletePostPage(w http.ResponseWriter, r *http.Request) {
	err := h.posts.deletePost(chi.URLParam(r, "id"), h.db)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "error while delete post")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// JSON decoder helper
func (h *Handler) decodePost(w http.ResponseWriter, r *http.Request) *Post {
	post := &Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		h.sendError(w, http.StatusInternalServerError, err, "error while decoding post body")
		return nil
	}
	return post
}
