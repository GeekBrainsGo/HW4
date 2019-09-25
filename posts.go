package main

import "database/sql"

// PostItem - объект поста в блоге
type PostItem struct {
	Id     	  int64 `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Labels    []string `json:"labels"`
}

// PostItemSlice - массив постов в блоге
type PostItemSlice []PostItem

// Create - создает задачу в БД
func (post *PostItem) Create(db *sql.DB) (PostItem, error) {
	row := db.QueryRow(
		"INSERT INTO PostItems (Title,  Text) VALUES (?, ?)",
		post.Title, post.Text,
	)

	newPost := PostItem{}
	if err := row.Scan(&post.Id, &post.Title, &post.Text); err != nil {
		return PostItem{}, err
	}

	return newPost, nil
}

// Delete - удалить объект из базы
func (post *PostItem) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM PostItems WHERE ID = ?",
		post.Id,
	)

	return err
}

// Update - обновляет объект в БД
func (post *PostItem) Update(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE PostItems SET Title = ?, Text = ? WHERE Id = ?",
		post.Title, post.Text, post.Id,
	)

	return err
}

// GetPostItem - получение поста
func GetPostItem(postId int64, db *sql.DB) (PostItem, error) {
	row := db.QueryRow(
		"SELECT Id, Title, Text FROM PostItems WHERE ?",
		postId)

	getPost := PostItem{}
	if err := row.Scan(&getPost.Id, &getPost.Title, &getPost.Text); err != nil {
		return PostItem{}, err
	}

	return getPost, nil
}

// GetAllPostItems - получение всех постов
func GetAllPostItems(db *sql.DB) (PostItemSlice, error) {
	rows, err := db.Query("SELECT Id, Title, Text FROM PostItems")
	if err != nil {
		return nil, err
	}
	posts := make(PostItemSlice, 0, 8)
	for rows.Next() {
		post := PostItem{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Text); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
