package main

import (
	gofpdf "./gofpdf"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"fmt"
	"time"
	"os"

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
		rStr, status := ascii.FontAscii(newInput, f)
		if status == 500 {
			e := fillerText{Text: "500"}
			js, _ := json.Marshal(e)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			fmt.Println("Internal server error")
		} else if status == 400 {
			e := fillerText{Text: "400"}
			js, _ := json.Marshal(e)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			fmt.Println("Bad request")
		} else {
			txt := fillerText{Text: rStr}
			js, _ := json.Marshal(txt)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}
}

func createFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		str := r.Form["text"][0]
		fmt.Println(str)
		format := r.Form["format"][0]
		font := r.Form["font"][0]
		currentTime := time.Now()
		str, _ = ascii.FontAscii(str, font)
		if format == ".pdf" {
			pdf := gofpdf.New("P", "mm", "A4", "")
			pdf.AddPage()
			pdf.CellFormat(190, 7, str, "0", 0, "CM", false, 0, "")
			pdf.OutputFileAndClose("keko")
			fmt.Println("created pdf")
		} else {
			file, err := os.Create(currentTime.String() + format)
			if err != nil {
				fmt.Println("Cannot create file")
				return 
			}
			file.WriteString(str)
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
	http.HandleFunc("/download", createFile)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}