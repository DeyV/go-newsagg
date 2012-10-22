package main

import (
	"html/template"
	// "io"
	// "code.google.com/p/goconf/conf"
	"./rss"
	"fmt"
	"net/http"
)

var List []*rss.Entry = []*rss.Entry{
	{Title: "1 tytuł", Id: "1", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
	{Title: "1 tytuł", Id: "2", Summary: rss.Text{Body: "2321321"}},
}

func getEntryList(cat string, limit int) []*rss.Entry {
	return List
}

// func Len(v []interface{}) int {
// 	return len(v)
// }

var templActionHome *template.Template = template.Must(
	template.ParseGlob("templates/home/*.html"))

func actionHome(w http.ResponseWriter, req *http.Request) {
	lay := getLayoutTemplates()
	wr := &HtmlContainer{}

	// templActionHome.Funcs(template.FuncMap{"len": Len})
	data := HtmlAssigner{
		"List": getEntryList("", 10),
		"Test": "Test",
	}

	templActionHome.Execute(wr, data)

	lay.New("title").Parse("Najnowsze wpisy ze świata PHP")

	lay.Execute(w, wr.getHtml())
}

func actionTest(w http.ResponseWriter, req *http.Request) {

	fmt.Fprint(w, "test")
}
