package models

// ErrorModel - Ошибка отвечаемая сервером
type ErrorModel struct {
	Code     int         `json:"code"`
	Err      string      `json:"error"`
	Desc     string      `json:"desc"`
	Internal interface{} `json:"internal"`
}
