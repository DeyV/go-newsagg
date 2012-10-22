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
	"time"
)

// 
var dbmap *gorp.DbMap
var db *sql.DB

var config *conf.ConfigFile
var debugShowTime bool

func init() {
	var err error
	config, err = conf.ReadConfigFile("config/project.ini")
	if err != nil {
		panic("Cant read config file: config/project.ini")
	}

	debugShowTime = config.GetBoolDef("debug", "showTime", true)

	// runtime.GOMAXPROCS(4) 

	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open(config.GetStringDef("db", "type", "sqlite3"), config.GetStringDef("db", "dns", "data/newsagg.db"))
	if err != nil {
		panic(err)
	}
	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
}

func main() {
	defer db.Close()

	RegModelsSchema(dbmap)

	runServer()
}

func runServer() {
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
		var t1 time.Time
		if debugShowTime {
			t1 = time.Now()
		}

		handler(w, req)

		if debugShowTime {
			t := time.Now().Sub(t1)
			fmt.Printf("Action %s, time: %s\n", pattern, t)
		}
	})
}
