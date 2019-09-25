package mysql

import (
	"blog/app/model"
)

// PostRepository ...
type PostRepository struct {
	mysql *Database
}

// Create ...
func (r *PostRepository) Create(p *model.Post) error {
	return r.mysql.db.QueryRow(
		"INSERT INTO posts (title, short, body) VALUES (?, ?, ?);",
		p.Title,
		p.Short,
		p.Body,
	).Scan(&p.ID)
}

// Update ...
func (r *PostRepository) Update(p *model.Post) error {
	_, err := r.mysql.db.Query(
		"UPDATE posts SET title = ?, short = ?, body = ? WHERE id = ?;",
		p.Title,
		p.Short,
		p.Body,
		p.ID,
	)
	return err
}

// Delete ...
func (r *PostRepository) Delete(id int64) error {
	_, err := r.mysql.db.Query(
		"DELETE FROM posts WHERE id = ?;",
		id,
	)
	return err
}

// Find ...
func (r *PostRepository) Find(id int64) (*model.Post, error) {
	p := &model.Post{}
	if err := r.mysql.db.QueryRow(
		"SELECT id, title, short, body, created_at, updated_at FROM posts WHERE id = ?;",
		id,
	).Scan(
		&p.ID,
		&p.Title,
		&p.Short,
		&p.Body,
		&p.Created,
		&p.Updated,
	); err != nil {
		return nil, err
	}

	return p, nil
}

// FindAll ...
func (r *PostRepository) FindAll() (*model.PostItemsSlice, error) {

	rows, err := r.mysql.db.Query("SELECT id, title, short, body, created_at, updated_at FROM posts;")
	if err != nil {
		return nil, err
	}
	posts := make(model.PostItemsSlice, 0, 8)
	for rows.Next() {
		post := model.Post{}
		if err := rows.Scan(&post.ID, &post.Title, &post.Short, &post.Body, &post.Created, &post.Updated); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}
