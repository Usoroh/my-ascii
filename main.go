package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	// "strings"
	"fmt"

	ascii "./ascii"
)

type fillerText struct {
	Text string
}

var wAdress *http.ResponseWriter


func serveTemplate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		wAdress = &w
		if r.URL.Path != "/" {
			fmt.Println(404)
			t, _ := template.ParseFiles("404.html")
			t.Execute(w, nil)
			return
		}
		// fmt.Println(w)
		t, _ := template.ParseFiles("index.html")
		fmt.Println(t)
		t.Execute(w, nil)
	}
}

func chooseFont(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	if r.Method == "GET" {

		str := r.Form["string"][0]
		f := r.Form["font"][0]
		// newInput := strings.Replace(str, "\n", "\n", -1)
		rStr, status := ascii.FontAscii(str, f)
		if status == 500 {
			e := fillerText{Text: "500"}
			js, _ := json.Marshal(e)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			fmt.Println("Internal server error")
		} else {
			txt := fillerText{Text: rStr}
			js, _ := json.Marshal(txt)
			fmt.Println(string(js))
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}
}

// func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 	w.WriteHeader(status)
// 	if status == 500 {
// 		fmt.Fprint(w, "Custom 500")
// 		t, _ := template.ParseFiles("500.html")
// 		t.Execute(w, nil)
// 	}

// 	// fmt.Println(name)
// 	// t, _ := template.ParseFiles(name+".html")
// 	// t.Execute(w, nil)
// }

func main() {
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/font", chooseFont)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
