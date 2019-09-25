package database

// Database ...
type Database interface {
	Post() PostRepository
}
