package server

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("Blog ID:", blogID)

	data, _ := ioutil.ReadAll(r.Body)

	blog := models.BlogItem{}
	_ = json.Unmarshal(data, &blog)
	blog.ID = blogID

	if err := blog.UpdateBlog(serv.db); err != nil {
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
	http.Redirect(w, r, "/", 301)
}
