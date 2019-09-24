/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 23.09.2019 19:33
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"time"
)

// Post is the base post type.
type Post struct {
	ID      string
	Title   string        `json:"title"`
	Date    string        `json:"date"` // todo convert to DateTime
	Summary string        `json:"summary"`
	Body    template.HTML `json:"body"`
}

//type dbPosts map[string]Post
type dbPosts []Post

// get one or all posts
func (p *dbPosts) getPosts(id string, db *sql.DB) (dbPosts, error) {
	var rows *sql.Rows
	var err error
	var posts = dbPosts{}
	if id != "" {
		rows, err = db.Query(GETONEPOST, id)
	} else {
		rows, err = db.Query(GETALLPOSTS)
	}
	if err != nil {
		return posts, fmt.Errorf("error in db.query: %v", err)
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Summary, &post.Body, &post.Date)
		if err != nil {
			return posts, fmt.Errorf("error in row.scan: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Close(); err != nil {
		return posts, err
	}
	return posts, nil
}

// create one post
func (p *dbPosts) createPost(post *Post, db *sql.DB) error {
	_, err := db.Exec(INSERTPOST, post.Title, post.Summary, post.Body)
	return err
}

// delete one post
func (p *dbPosts) deletePost(id string, db *sql.DB) error {
	delTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(DELETEPOST, delTime, id)
	return err
}

// update one post
func (p *dbPosts) updatePost(post *Post, db *sql.DB) error {
	_, err := db.Exec(UPDATEPOST, post.Title, post.Summary, post.Body, post.ID)
	return err
}
