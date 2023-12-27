package render

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

//render templates

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles(tmpl)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
