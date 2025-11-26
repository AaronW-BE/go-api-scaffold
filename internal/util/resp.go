package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	resp := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func Text(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
		return
	}
}
