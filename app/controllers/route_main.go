package controllers

import (
	"html/template"
	"net/http"
)

//Serverがスタートしたら、Topページが立ち上がる処理
func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/templates/top.html")
	t.Execute(w, nil)
}
