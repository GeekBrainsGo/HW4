package models

// Page - страница доступная шаблонизатору
type Page struct {
	Title   string
	Command string
	Data    interface{}
	// Title, Content string
	// Blogs          BlogItemSlice
}
