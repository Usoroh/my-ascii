package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"fmt"

	ascii "./ascii"
)

type fillerText struct {
	Text string
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			fmt.Println(404)
			t, _ := template.ParseFiles("404.html")
			t.Execute(w, nil)
			return
		}
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
}

func chooseFont(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()


	if r.Method == "GET" {

		str := r.Form["string"][0]
		f := r.Form["font"][0]

		newInput := strings.Replace(str, "\r\n", "\\n", -1)
		// args := []string{str, f}
		// txt := fillerText{Text: Font(args)}
		rStr, status := ascii.FontAscii(newInput, f)
		if status == 500 {
			t, err := template.ParseFiles("500.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
			fmt.Println("Internal server error")
			return
		} else {
			txt := fillerText{Text: rStr}
			js, _ := json.Marshal(txt)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}
}


func main() {
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/font", chooseFont)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
