package categorycontroller

import (
	"crud/src/entities"
	"crud/src/models/categorymodel"
	"html/template"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("src/views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmplt, err := template.ParseFiles("src/views/category/create.html")

		if err != nil {
			panic(err)
		}

		tmplt.Execute(w, nil)
	}

	if r.Method  == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			tmplt, _ := template.ParseFiles("src/views/category/create.html")
			tmplt.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
