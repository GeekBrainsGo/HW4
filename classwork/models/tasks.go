package models

import "database/sql"

// TaskItem - объект задачи
type TaskItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

// TaskItemSlice - массив задач
type TaskItemSlice []TaskItem

// Insert - добавляет задачу в БД
func (task *TaskItem) Insert(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO TaskItems (ID, Text, Completed) VALUES (?, ?, ?)",
		task.ID, task.Text, task.Completed,
	)
	return err
}

// Delete - удалят объект из базы
func (task *TaskItem) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM TaskItems WHERE ID = ?",
		task.ID,
	)
	return err
}

// Update - обновляет объект в БД
func (task *TaskItem) Update(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE TaskItems SET Text = ?, Completed = ? WHERE ID = ?",
		task.Text, task.Completed, task.ID,
	)
	return err
}

// GetAllTaskItems - получение всех задач
func GetAllTaskItems(db *sql.DB) (TaskItemSlice, error) {
	rows, err := db.Query("SELECT ID, Text, Completed FROM TaskItems")
	if err != nil {
		return nil, err
	}
	tasks := make(TaskItemSlice, 0, 8)
	for rows.Next() {
		task := TaskItem{}
		if err := rows.Scan(&task.ID, &task.Text, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}
