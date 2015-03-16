package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

var templates map[string]*template.Template

func main() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob("./layouts/*")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob("./includes/*")
	if err != nil {
		panic(err)
	}

	for _, layout := range layouts {
		files := append(includes, layout)
		templates[layout] = template.Must(template.ParseFiles(files))
	}

	fmt.Println(templates("base"))

}
