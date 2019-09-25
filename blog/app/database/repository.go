package database

import "blog/app/model"

// PostRepository ...
type PostRepository interface {
	Create(*model.Post) error
	Update(*model.Post) error
	Delete(int64) error
	Find(int64) (*model.Post, error)
	FindAll() (*model.PostItemsSlice, error)
}
