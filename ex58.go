package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}
func (S String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello string!!!\n\n")
	/*
	r.ParseForm()
	fmt.Println("Form : ",r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("param : ", r.Form["test_param"])
	*/
	path, _ := os.Getwd()
	path += "/ex58.go"
	fmt.Fprint(w, path)
	fmt.Fprint(w, "\n\n")

	content, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	w.Write(content)
}
func (S *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello struct!!\n\n")
	path, _ := os.Getwd()
	path += "/ex57.go"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	w.Write(content)
}
func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers"})
	http.ListenAndServe(":4000", nil)
}
