package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/ftc", handler)
	http.Handle("/", http.FileServer(http.Dir("C:/Users/Yousuf/Desktop/first-scouting/templateGUI")))
	panic(http.ListenAndServe(":4278", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("template.html")
	w.Header().Set("Content-Type", "text/html")

	//fmt.Fprintf(w, bytes.NewBuffer(f).String())
	io.WriteString(w, bytes.NewBuffer(f).String())
}
