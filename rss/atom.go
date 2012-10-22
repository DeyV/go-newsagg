package rss

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	XMLName xml.Name  `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Person    `xml:"author"`
	Entry   []*Entry  `xml:"entry"`
}

type Entry struct {
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
	Author  Person    `xml:"author"`
	Summary Text      `xml:"summary"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri"`
	Email    string `xml:"email"`
	InnerXML string `xml:",innerxml"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

func ParseTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Now() // zwraca NOW zamiast panic
	}
	return t
}

func NewText(text string) Text {
	return Text{
		Body: text,
	}
}

func ReadAtomXml(data []byte) (*Feed, error) { /* r io.Reader */
	feed := &Feed{}

	err := xml.Unmarshal(data, feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func WriteAtomXml(f *Feed) ([]byte, error) {
	return xml.Marshal(f)
}
