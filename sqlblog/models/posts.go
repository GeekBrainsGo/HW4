package models

import (
	"database/sql"
	"html/template"
	"time"
)

// Post stands for post object.
type Post struct {
	ID      string        `json:"ID"`
	Title   string        `json:"Title"`
	Author  string        `json:"Author"`
	Created string        `json:"Created"`
	Content template.HTML `json:"Content"`
}

// Posts stands for array of posts.
type Posts []Post

// Insert post to database.
func (p *Post) Insert(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO Posts (Title, Author, Content) VALUES (?, ?, ?)",
		p.Title, p.Author, p.Content,
	)
	return err
}

// Delete deletes post from database.
func (p *Post) Delete(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM Posts WHERE ID = ?",
		p.ID,
	)
	return err
}

// Update updates post in database.
func (p *Post) Update(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE Post SET Title = ?, Author = ?, Created = ?, Content = ? WHERE ID = ?",
		p.Title, p.Author, time.Now().Format("15:04:05 02-01-2006"), p.Content, p.ID,
	)
	return err
}

// AllPosts return all posts from database.
func AllPosts(db *sql.DB) (Posts, error) {
	rows, err := db.Query("SELECT ID, Title, Author, Created, Content FROM Posts")
	if err != nil {
		return nil, err
	}
	posts := make(Posts, 0, 8)
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.Title, &post.Author, &post.Created, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}
