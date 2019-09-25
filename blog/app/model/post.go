package model

// Post ...
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Short   string `json:"short"`
	Body    string `json:"body"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

// PostItemsSlice ...
type PostItemsSlice []Post
