/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 23.09.2019 21:33
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

// Constants
const (
	SERVADDR     = ":8080"
	TEMPLATEEXT  = "*.gohtml"
	TEMPLATEPATH = "templates"
	POSTSURL     = "/posts"
	EDITURL      = "/edit"
	CREATEURL    = "/create"
	APIURL       = "/api/v1"
	STATICPATH   = "/static"
	DBNAME       = "blog"
	DSN          = "/" + DBNAME + "?charset=utf8&interpolateParams=true"
	TABLENAME    = "posts"
	GETALLPOSTS  = "SELECT id, title, summary, body, DATE_FORMAT(updated, '%d.%m.%Y %H:%i') FROM " + TABLENAME + " WHERE deleted IS NULL ORDER BY id DESC"
	GETONEPOST   = "SELECT id, title, summary, body, DATE_FORMAT(updated, '%d.%m.%Y %H:%i') FROM " + TABLENAME + " WHERE deleted IS NULL AND id = ?"
	//DELETEPOST   = "DELETE FROM " + TABLENAME + " WHERE id = ?"
	DELETEPOST = "UPDATE " + TABLENAME + " SET deleted = ? WHERE id = ?"
	INSERTPOST = "INSERT INTO " + TABLENAME + " (title, summary, body) VALUES(?, ?, ?)"
	UPDATEPOST = "UPDATE " + TABLENAME + " SET title = ?, summary = ?, body = ? WHERE ID = ?"
)
