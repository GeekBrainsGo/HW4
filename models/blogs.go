package models

import (
	"database/sql"
)

// BlogItem - объект блога
type BlogItem struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// BlogItemSlice - массив блогов
type BlogItemSlice []BlogItem

// GetAllBlogItems - получение всех блогов
func GetAllBlogItems(db *sql.DB) (BlogItemSlice, error) {

	rows, err := db.Query("SELECT ID, Title FROM BlogItems")
	if err != nil {
		return nil, err
	}
	blogs := make(BlogItemSlice, 0, 8)
	for rows.Next() {
		blog := BlogItem{}
		if err := rows.Scan(&blog.ID, &blog.Title); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, err
}

// Delete - удалят объект из базы
func (blog *BlogItem) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM blogitems WHERE ID = ?",
		blog.ID,
	)
	return err
}

// Insert - добавляет задачу в БД
// func (blog *BlogItem) Insert(db *sql.DB) error {
// 	_, err := db.Exec(
// 		"INSERT INTO BlogItems (ID, Text) VALUES (?, ?)",
// 		blog.ID, blog.Title,
// 	)
// 	return err
// }

// Update - обновляет объект в БД
// func (task *TaskItem) Update(db *sql.DB) error {
// 	_, err := db.Exec(
// 		"UPDATE TaskItems SET Text = ?, Completed = ? WHERE ID = ?",
// 		task.Text, task.Completed, task.ID,
// 	)
// 	return err
// }
