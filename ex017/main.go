package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// Servidor web sencillo
func main() {
	http.HandleFunc("/", home)
	log.Println("PORT 8081 LISTENING")
	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	path, _ := filepath.Abs("./cursoGo/ex017/web/index.html")
	http.ServeFile(w, r, path)
	log.Println("INDEX")
}
