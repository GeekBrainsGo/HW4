/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 23.09.2019 20:33
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// routes preparer
func (h *Handler) prepareRoutes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.HandleFunc("/", h.mainPageForm)
	mux.Route(APIURL+POSTSURL, func(r chi.Router) {
		r.Post(CREATEURL, h.createPostPage)
		r.Put("/{id}", h.editPostPage)
		r.Delete("/{id}", h.deletePostPage)
	})
	mux.Route(POSTSURL, func(r chi.Router) {
		r.Get("/", h.postsPageForm)
		r.Get(CREATEURL, h.createPostPageForm)
		r.Get(EDITURL+"/", h.editPostPageForm)
	})
	mux.Handle(STATICPATH+"/*", http.StripPrefix(STATICPATH, http.FileServer(http.Dir("."+STATICPATH))))
	return mux
}
