package category

import (
	"fmt"
	"github.com/coopernurse/gorp"
)

const (
	MinFontSize = 8
	MaxFontSize = 18
)

var categories []*Category
var err error

func getExamplCat() []interface{} {
	tab := []interface{}{
		&Category{Id: 1, Name: "PHP", Count: 15, LastUpdate: "2012-09-09T13:42:23"},
		&Category{Id: 2, Name: "SQL", Count: 5, LastUpdate: "2012-09-09T13:42:23"},
		&Category{Id: 3, Name: "Symfony", Count: 7, LastUpdate: "2012-09-09T13:42:23"},
		&Category{Id: 4, Name: "Zend", Count: 3, LastUpdate: "2012-09-09T13:42:23"},
		&Category{Id: 5, Name: "PostgreSQL", Count: 2, LastUpdate: "2012-09-09T13:42:23"},
	}
	return tab
}

func init() {
	// tab := getAllCats(dbmap)
	tab := getExamplCat()
	initData(tab)
}

func InitDb(dbmap *gorp.DbMap) {
	// tab := getAllCats(dbmap)
	tab := getExamplCat()
	initData(tab)
}

func initData(tab []interface{}) {
	count := len(tab)
	categories = make([]*Category, count)

	var max int

	for key, row := range tab {
		cat := row.(*Category)
		categories[key] = cat

		if max < cat.Count {
			max = cat.Count
		}
	}

	for _, cat := range categories {
		cat.FontSize = getFontSize(max, cat.Count)
	}
}

func getAllCats(dbmap *gorp.DbMap) []interface{} {
	tab, err := dbmap.Select(Category{}, "SELECT * FROM Category")
	if err != nil {
		fmt.Println("Problem z Category InitDb")
	}
	return tab
}

var zakres float64 = MaxFontSize - MinFontSize

func getFontSize(max, val int) int {
	proc := float64(val) / float64(max)
	return MinFontSize + int(proc*zakres)
}

func GetTagCloud() []*Category {
	return categories
}
