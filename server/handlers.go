package server

import (
	"html/template"
	"io/ioutil"
	"myblog/models"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/go-chi/chi"
)

// getTemplateHandler - возвращает шаблон
func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateName := chi.URLParam(r, "template")

	if templateName == "" {
		templateName = serv.indexTemplate
	}

	file, err := os.Open(path.Join(serv.rootDir, serv.templatesDir, templateName))
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		serv.SendInternalErr(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	// templ, err := template.New("Page").Parse(string(data))
	templ := template.Must(template.New("page").Parse(string(data)))

	blogs, err := models.GetAllBlogItems(serv.db)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Blogs = blogs

	if err := templ.Execute(w, serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// deleteBlogHandler - удаляет блог
func (serv *Server) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	bl := chi.URLParam(r, "id")
	blogID, _ := strconv.ParseInt(bl, 10, 64)

	blog := models.BlogItem{ID: blogID}

	if err := blog.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}
