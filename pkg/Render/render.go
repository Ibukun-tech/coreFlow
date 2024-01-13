package Render

import (
	"fmt"
	"net/http"
	"text/template"
)

func ParseTemplate(w http.ResponseWriter, locat string) {

	parsedFile, _ := template.ParseFiles("./template/" + locat)
	err := parsedFile.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
