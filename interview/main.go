package main

import (
	"fmt"
	"github.com/joecomscience/go-projects/interview/db"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	db.ConnectDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveTemplate)
	mux.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/index.html")
	t.Execute(w, map[string]interface{}{
		"auth": os.Getenv("AUTH"),
	})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*8)
	err := r.ParseMultipartForm(8 << 20)
	if err != nil {
		responseError(w, fmt.Sprint("error max file size"))
		return
	}

	if r.FormValue("auth") != os.Getenv("AUTH") {
		responseError(w, fmt.Sprint("permission denied"))
		return
	}

	f, h, err := r.FormFile("image_file")
	if err != nil {
		responseError(w, fmt.Sprintf("error recive file: %v\n", err))
		return
	}
	defer f.Close()

	file := strings.Split(h.Filename, ".")
	filename := fmt.Sprintf("%s-*.%s", file[0], file[1])
	t, err := ioutil.TempFile("images", filename)
	if err != nil {
		responseError(w, fmt.Sprintf("error create temp file: %v\n", err))
		return
	}
	defer t.Close()

	fb, err := ioutil.ReadAll(f)
	if err != nil {
		responseError(w, fmt.Sprintf("error create temp file: %v\n", err))
		return
	}
	t.Write(fb)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	i := db.Image{
		Size:     h.Size,
		Filename: t.Name(),
	}

	for hd, _ := range h.Header {
		i.Header = append(i.Header, hd)
	}

	if err := i.Insert(); err != nil {
		responseError(w, fmt.Sprintf("save content to db error: %v\n", err))
		return
	}
}

func responseError(w http.ResponseWriter, msg string) {
	fmt.Printf("%s\n", msg)
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintln(w, msg)
}
