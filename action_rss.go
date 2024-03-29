package main

import (
	"github.com/DeyV/go-newsagg/rss"
	//	"code.google.com/p/goconf/conf"

	"net/http"
)

func init() {
	HandleFunc("/rss", actionRssAtom)
	HandleFunc("/rss/atom", actionRssAtom)
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
