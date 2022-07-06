package controllers

import (
	"net/http"
)

//Serverがスタートしたら、Topページが立ち上がる処理
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "Top")
}
