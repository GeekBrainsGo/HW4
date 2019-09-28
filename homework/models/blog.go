package models

import (
	"database/sql"
)

type BlogItems []BlogItem

type BlogItem struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Contents string `json:"contents,omitempty"`
}

//GetSingleBlogItem - get a blogItem from DB
func GetSingleBlogItem(id string, db *sql.DB) (BlogItem, error) {
	blogItem := BlogItem{}
	rows, err := db.Query("SELECT ID, Title, Contents FROM blogs where ID=?", id)
	if err != nil {
		return blogItem, err
	}
	if rows.Next() {
		if err := rows.Scan(&blogItem.ID, &blogItem.Title, &blogItem.Contents); err != nil {
			return blogItem, err
		}
	}
	return blogItem, err
}

//GetAllBlogItems - get all blogItems from DB
func GetAllBlogItems(db *sql.DB) ([]BlogItem, error) {
	blogItem := BlogItem{}
	var blogItems BlogItems
	rows, err := db.Query("SELECT ID, Title, Contents FROM blogs")
	if err != nil {
		return blogItems, err
	}
	for rows.Next() {
		if err := rows.Scan(&blogItem.ID, &blogItem.Title, &blogItem.Contents); err != nil {
			return blogItems, err
		}
		blogItems = append(blogItems, blogItem)
	}
	return blogItems, err
}

//Update - updates blogitem in the DB
func (blogItem BlogItem) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE blogs SET Title=?, Contents=? WHERE ID=?)", blogItem.Title, blogItem.Contents, blogItem.ID)
	if err != nil {
		return err
	}
	return nil
}

//Insert - inserts blogitem into the DB
func (blogItem BlogItem) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO blogs (ID, Title, Contents) VALUES (?,?,?)", blogItem.ID, blogItem.Title, blogItem.Contents)
	if err != nil {
		return err
	}
	return nil
}

//Delete - deletes blogitem from the DB by ID
func (blogItem BlogItem) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM blogs WHERE ID=?", blogItem.ID)
	if err != nil {
		return err
	}
	return nil
}
