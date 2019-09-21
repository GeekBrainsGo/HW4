package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"serv/models"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
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

	templ, err := template.New("Page").Parse(string(data))
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	tasks, err := models.GetAllTaskItems(serv.db)
	if err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	serv.Page.Tasks = tasks

	if err := templ.Execute(w, serv.Page); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// postTaskHandler - добавляет новую заявку
func (serv *Server) postTaskHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	task := models.TaskItem{}
	_ = json.Unmarshal(data, &task)

	task.ID = uuid.NewV4().String()

	if err := task.Insert(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}

	data, _ = json.Marshal(task)
	w.Write(data)
}

// deleteTaskHandler - удаляет задачу
func (serv *Server) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	task := models.TaskItem{ID: taskID}
	if err := task.Delete(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}

// putTaskHandler - обновляет задачу
func (serv *Server) putTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	data, _ := ioutil.ReadAll(r.Body)

	task := models.TaskItem{}
	_ = json.Unmarshal(data, &task)
	task.ID = taskID

	if err := task.Update(serv.db); err != nil {
		serv.SendInternalErr(w, err)
		return
	}
}
