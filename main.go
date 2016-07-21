package main  

import (
	"net/http"
	"os"
	"text/template"
	"io"
	"fmt"
	"time"
)

const STATIC_URL string = "/public/"
const STATIC_ROOT string = "public/"

func populateTemplate() *template.Template  {
	result := template.New("templates")

	basePath := "views"

	templateFolder,_ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathRaw,_ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _,pathInfo := range templatePathRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	fmt.Println("Static file: ", static_file)

	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err==nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return 
		}
	}
	http.NotFound(w, req)
}

func main() {
	templates := populateTemplate()
	http.HandleFunc("/",
		func (w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")
			fmt.Println("template:", template)

			if template != nil {
				template.Execute(w, nil)
			}else{
				w.WriteHeader(404)
			}
		})

	http.HandleFunc(STATIC_URL, StaticHandler)

	//http.HandleFunc("/index", func(w))

	http.ListenAndServeTLS(":80", nil)
}
