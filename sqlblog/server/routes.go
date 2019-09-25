package server

import "github.com/go-chi/chi"

func (s *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{template}", s.templateHandle)
		r.Route("/api/v1", func(r chi.Router) {
			r.Post("/posts", s.postHandle)
			r.Delete("/posts/{id}", s.deleteHandle)
			r.Put("/posts/{id}", s.putHandle)
		})
	})
}
