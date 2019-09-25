package mysql

import (
	"blog/app/database"
	"database/sql"
)

// Database ...
type Database struct {
	db             *sql.DB
	postRepository *PostRepository
}

// New ...
func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

// Post ...
func (db *Database) Post() database.PostRepository {
	if db.postRepository != nil {
		return db.postRepository
	}

	db.postRepository = &PostRepository{
		mysql: db,
	}

	return db.postRepository
}
