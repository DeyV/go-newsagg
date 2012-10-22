package main

import (
	"html/template"
)

type HtmlContainer struct {
	body []byte
}

func (h *HtmlContainer) Write(b []byte) (int, error) {
	h.body = append(h.body, b...)
	return len(b), nil
}

func (h *HtmlContainer) getHtml() template.HTML {
	return template.HTML(string(h.body))
}

type HtmlAssigner map[string]interface{}

func getLayoutTemplates() *template.Template {
	var templLayouts *template.Template = template.Must(
		template.ParseGlob("templates/layouts/*.html"))

	temp, _ := templLayouts.Clone()

	return temp 
}
