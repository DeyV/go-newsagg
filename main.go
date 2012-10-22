package main

import (
	//	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/DeyV/go-sqlite3-win64"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/coopernurse/gorp"

	"./conf"
	"net/http"
	// "runtime"

	// "time"
)

// 
var dbmap *gorp.DbMap
var db *sql.DB

var config *conf.ConfigFile

func init() {
	config, _ = conf.ReadConfigFile("config/project.ini")

	// runtime.GOMAXPROCS(4) 

	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./data/Planeta.sqlite")
	if err != nil {
		panic(err)
	}
	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
}

func main() {
	defer db.Close()

	runServer()
}

func runServer() {
	// http.HandleFunc("/form", actionForm)
	// http.HandleFunc("/tags", actionTags)
	HandleFunc("/", actionHome)
	HandleFunc("/test", actionTest)
	HandleFunc("/rss", actionRssAtom)
	HandleFunc("/rss/atom", actionRssAtom)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	siteUrl, _ := config.GetString("default", "siteUrl")

	println("start ", siteUrl)
	if err := http.ListenAndServe(siteUrl, nil); err != nil {
		panic(err)
	}
	fmt.Println("end")
}

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		// t1 := time.Now()

		handler(w, req)

		// t := time.Now().Sub(t1)
		// fmt.Printf("Action %s, time: %s\n", pattern, t)
	})
}
