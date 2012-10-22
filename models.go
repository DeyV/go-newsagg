package main

import (
	"./category"
	"./rss"
	"encoding/xml"
	"github.com/coopernurse/gorp"
)

func RegModelsSchema(dbmap *gorp.DbMap) {
	category.RegSchema(dbmap)
	category.InitDb(dbmap)
}

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

type RssConfig struct {
	Title,
	HrefRel,
	HrefSelf,
	Author,
	AuthorXML,
	Updated string
}

func newAtom(entry []*rss.Entry, conf RssConfig) *rss.Feed {

	var atomFeed = &rss.Feed{
		XMLName: xml.Name{"http://www.w3.org/2005/Atom", "feed"},
		Title:   conf.Title,
		Link: []rss.Link{
			{Rel: "alternate", Href: conf.HrefRel},
			{Rel: "self", Href: conf.HrefSelf},
		},
		Id:      conf.HrefRel,
		Updated: rss.ParseTime(conf.Updated),
		Author: rss.Person{
			Name:     conf.Author,
			InnerXML: conf.AuthorXML,
		},
		Entry: entry,
	}
	return atomFeed
}
