/*
 * HomeWork-4: Simple blog - MySQL
 * Created on 23.09.2019 20:22
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error model.
type Error struct {
	ErrCode  int    `json:"code"`
	ErrText  string `json:"error"`
	ErrDescr string `json:"descr"`
}

// errors helper
func (e *Error) sendError(w http.ResponseWriter, code int, err error, descr string) {
	log.Println(descr, "-", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errMsg := Error{
		ErrCode:  code,
		ErrText:  err.Error(),
		ErrDescr: descr,
	}
	data, err := json.Marshal(errMsg)
	if err != nil {
		log.Println("Can't marshal error data:", err)
		return
	}
	if _, err = w.Write(data); err != nil {
		log.Println("Can't write to ResponseWriter:", err)
	}
}
