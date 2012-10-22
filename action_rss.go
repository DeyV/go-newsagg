package main

import (
	"./rss"
	//	"code.google.com/p/goconf/conf"
	"encoding/xml"
	"net/http"
)

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

func actionRssAtom(w http.ResponseWriter, req *http.Request) {
	cat := req.FormValue("cat")
	title := ""

	if cat != "" {
		title = "[" + cat + "]"
	}

	rssConf := RssConfig{
		Title:    config.GetStringDef("rss", "title", "Planeta RSS") + " " + title,
		HrefRel:  config.GetStringDef("page", "link", ""),
		HrefSelf: config.GetStringDef("rss", "link", ""),
		Author:   config.GetStringDef("rss", "author", "PHP.pl"),
		Updated:  "2012-09-21T19:51:50+02:00",
	}

	xmldata, _ := rss.WriteAtomXml(newAtom(getEntryList(cat, 10), rssConf))

	w.Header().Add("content-type", "text/atom+xml") // application
	w.Write(xmldata)
}
