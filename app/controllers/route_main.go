package controllers

import (
	"log"
	"net/http"
)

//Serverがスタートしたら、Topページが立ち上がる処理
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "Top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
