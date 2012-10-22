package category

import (
	"github.com/coopernurse/gorp"
)

type Category struct {
	Id         int
	Name       string
	Count      int
	LastUpdate string
	FontSize   int `db:-`
}

type EntryInCategory struct {
	Id         int
	CategoryId int
	EntryId    int
}

func RegSchema(dbmap *gorp.DbMap) {
	dbmap.AddTable(Category{}).SetKeys(true, "Id")
	dbmap.AddTable(EntryInCategory{}).SetKeys(true, "Id")

	dbmap.CreateTables()
}
