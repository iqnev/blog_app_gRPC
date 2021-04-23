package rest

import "net/http"

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ivelin"))
}
