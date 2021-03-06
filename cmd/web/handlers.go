package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}


	err = ts.Execute(w, nil)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id<1{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed\n", 405)
	}
	w.Write([]byte("Creating a snippet..."))
}


