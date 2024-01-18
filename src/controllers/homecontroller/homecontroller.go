package homecontroller

import (
	"text/template"
	"net/http"
)


func Welcome(write  http.ResponseWriter, reader *http.Request) {
	tmplt, err := template.ParseFiles("src/views/home/index.html")

	if err != nil {
		panic(err)
	}

	tmplt.Execute(write, nil)
}