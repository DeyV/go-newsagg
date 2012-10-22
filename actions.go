package main

import (
	"html/template"
	// "io"
	// "code.google.com/p/goconf/conf"
	"./category"
	"fmt"
	"net/http"
)

func init() {
	HandleFunc("/", actionHome)
	HandleFunc("/test", actionTest)
	HandleFunc("/k", actionCategories)
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

	lay.New("title").Parse(config.GetStringDef("page", "title", "Page Title"))

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
	lay.New("title").Parse(config.GetStringDef("page", "title", "Page Categories"))

	err = lay.Execute(w, wr.getHtml())
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func actionTest(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "test")
}
