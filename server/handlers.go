package server

import (
	"encoding/json"
	"io/ioutil"
	"myblog/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// getTemplateHandler - возвращает шаблон
func (serv *Server) getTemplateHandler(w http.ResponseWriter, r *http.Request) {

	blogs, err := models.GetAllBlogItems(serv.db)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Title = "Мой блог"
	serv.Page.Data = blogs
	serv.Page.Command = "index"

	if err := serv.dictionary["BLOGS"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// viewBlogHandler - просмотр блога
func (serv *Server) viewBlogHandler(w http.ResponseWriter, r *http.Request) {

	blogIDStr := chi.URLParam(r, "id")
	blogID, _ := strconv.ParseInt(blogIDStr, 10, 64)

	blog, err := models.GetBlogItem(serv.db, blogID)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Title = "Статья"
	serv.Page.Data = blog
	serv.Page.Command = "view"

	if err := serv.dictionary["BLOG"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// editBlogHandler - редактирование блога
func (serv *Server) editBlogHandler(w http.ResponseWriter, r *http.Request) {

	bl := chi.URLParam(r, "id")
	blogID, _ := strconv.ParseInt(bl, 10, 64)

	blog, err := models.GetBlogItem(serv.db, blogID)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Title = "Редактирование"
	serv.Page.Data = blog
	serv.Page.Command = "edit"

	if err := serv.dictionary["BLOG"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// putBlogHandler - обновляет блог
func (serv *Server) putBlogHandler(w http.ResponseWriter, r *http.Request) {
	bl := chi.URLParam(r, "id")
	blogID, _ := strconv.ParseInt(bl, 10, 64)

	data, _ := ioutil.ReadAll(r.Body)

	blog := models.BlogItem{}
	_ = json.Unmarshal(data, &blog)
	blog.ID = blogID

	if err := blog.UpdateBlog(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}

// addGetBlogHandler - добавление блога
func (serv *Server) addGetBlogHandler(w http.ResponseWriter, r *http.Request) {

	serv.Page.Title = "Добавление блога"
	serv.Page.Data = models.BlogItem{}
	serv.Page.Command = "new"

	if err := serv.dictionary["BLOG"].ExecuteTemplate(w, "base", serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// addBlogHandler - добавляет блог
func (serv *Server) addBlogHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	blog := models.BlogItem{}
	_ = json.Unmarshal(data, &blog)

	if err := blog.AddBlog(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
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
	http.Redirect(w, r, "/", 301)
}
