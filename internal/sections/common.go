package sections

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseTemplate(page string) (*template.Template, error) {
	partials := []string{
		"./internal/templates/base.layout.gohtml",
		"./internal/templates/header.partial.gohtml",
		"./internal/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./internal/templates/%s", page))

	for _, t := range partials {
		templateSlice = append(templateSlice, t)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil

}

func ReadJSON[T any](filename string) T {
	var item T

	data, err := ioutil.ReadFile(fmt.Sprintf("./internal/sections/data/%s", filename))
	if err != nil {
		log.Panic(err)
		return item
	}

	err = json.Unmarshal(data, &item)
	if err != nil {
		log.Panic(err)
		return item
	}

	log.Printf("File: %s has been read successfully", filename)
	return item
}

func RenderWithData(w http.ResponseWriter, page string, data interface{}) {
	tmpl, err := ParseTemplate(page)
	if err != nil {
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
