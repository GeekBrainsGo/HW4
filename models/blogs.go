package models

import (
	"database/sql"
	"errors"
)

// BlogItem - объект блога
type BlogItem struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"article"`
}

// BlogItemSlice - массив блогов
type BlogItemSlice []BlogItem

// AddBlog - обновляет объект в БД
func (blog *BlogItem) AddBlog(db *sql.DB) error {

	_, err := db.Exec(
		"INSERT INTO BlogItems (Title, Body) VALUES ( ?,  ? )",
		blog.Title, blog.Body,
	)
	return err
}

// UpdateBlog - обновляет объект в БД
func (blog *BlogItem) UpdateBlog(db *sql.DB) error {

	_, err := db.Exec(
		"UPDATE BlogItems SET Title = ?, Body = ? WHERE ID = ?",
		blog.Title, blog.Body, blog.ID,
	)
	return err
}

// GetAllBlogItems - получение всех блогов
func GetAllBlogItems(db *sql.DB) (BlogItemSlice, error) {

	rows, err := db.Query("SELECT ID, Title, Body FROM BlogItems")
	if err != nil {
		return nil, err
	}
	blogs := make(BlogItemSlice, 0, 8)
	for rows.Next() {
		blog := BlogItem{}
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Body); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, err
}

// GetAllBlogItems - получение всех блогов
func GetBlogItem(db *sql.DB, id int64) (BlogItem, error) {

	blog := BlogItem{}
	var err error

	row := db.QueryRow("SELECT ID, Title, Body FROM BlogItems WHERE ID = ?", id)

	if row == nil {
		return blog, errors.New("Пустое значение!")
	}

	if row != nil {
		if err := row.Scan(&blog.ID, &blog.Title, &blog.Body); err != nil {
			return blog, err
		}
	}
	return blog, err
}

// Delete - удалят объект из базы
func (blog *BlogItem) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM blogitems WHERE ID = ?",
		blog.ID,
	)
	return err
}

// Insert - добавляет блог в БД
func (blog *BlogItem) Insert(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO BlogItems (ID, Title) VALUES (?, ?)",
		blog.ID, blog.Title,
	)
	return err
}

// Update - обновляет объект в БД
// func (task *TaskItem) Update(db *sql.DB) error {
// 	_, err := db.Exec(
// 		"UPDATE TaskItems SET Text = ?, Completed = ? WHERE ID = ?",
// 		task.Text, task.Completed, task.ID,
// 	)
// 	return err
// }
