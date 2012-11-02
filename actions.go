package main

import (
	"html/template"
	// "io"
	// "code.google.com/p/goconf/conf"
	"fmt"
	// "github.com/gorilla/mux"
	"github.com/DeyV/go-newsagg/category"
	"net/http"
)

func init() {
	HandleFunc("/", actionHome)
	HandleFunc("/test", actionTest)
	// HandleFunc("/k", actionCategories)
	HandleFunc("/k", actionCategoryList)
	HandleFunc("/k/latest", actionCategories)
}

func actionHome(w http.ResponseWriter, req *http.Request) {

	var templActionHome *template.Template = template.Must(
		template.ParseGlob("templates/home/*.html"))

	lay := getLayoutTemplates()
	wr := &HtmlContainer{}

	// templActionHome.Funcs(template.FuncMap{"len": Len})
	data := HtmlAssigner{
		"List": getEntryList("", 10),
		"Test": "Test",
	}

	err := templActionHome.Execute(wr, data)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	lay.New("title").Parse("Najnowsze wpisy - " + config.GetStringDef("page", "title", ""))

	err = lay.Execute(w, wr.getHtml())
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func actionCategoryList(w http.ResponseWriter, req *http.Request) {

	Id := req.URL.Query().Get("id")
	fmt.Println(Id)

	// // param := req.Vars()
	// catId, ok := param["id"]
	// if ok == nil {
	// 	catId = 0
	// }

	var templActionHome *template.Template = template.Must(
		template.ParseGlob("templates/home/*.html"))

	lay := getLayoutTemplates()
	wr := &HtmlContainer{}

	// templActionHome.Funcs(template.FuncMap{"len": Len})
	data := HtmlAssigner{
		"List": getEntryList("", 10),
		"Test": "Test",
	}

	err := templActionHome.Execute(wr, data)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	lay.New("title").Parse("Najnowsze wpisy - " + config.GetStringDef("page", "title", ""))

	err = lay.Execute(w, wr.getHtml())
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func actionCategories(w http.ResponseWriter, req *http.Request) {
	var templActionCategories *template.Template = template.Must(
		template.ParseGlob("templates/categories/*.html"))

	wr := &HtmlContainer{}

	// templActionHome.Funcs(template.FuncMap{"len": Len})
	data := HtmlAssigner{
		"List": category.GetTagCloud(),
		"Test": "Test",
	}

	err := templActionCategories.Execute(wr, data)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	lay := getLayoutTemplates()
	lay.New("title").Parse("Lista Kategorii - " + config.GetStringDef("page", "title", ""))

	err = lay.Execute(w, wr.getHtml())
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func actionTest(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "test")
}
