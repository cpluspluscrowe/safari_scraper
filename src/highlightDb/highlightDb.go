package highlightDb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	Text string
	Id   []uint8
}

func InsertHighlights(highlights []string) {
	db := getDatabaseDriver()
	defer db.Close()
	createHighlightTable(db)
	for _, highlight := range highlights {
		insertHighlight(db, highlight)
	}
}

func getDatabaseDriver() *sql.DB {
	db, err := sql.Open("sqlite3", "./highlights.db")
	if err != nil {
		panic(err)
	}
	checkErr(err)
	return db
}

func GetHighlights() []Highlight {
	db, err := sql.Open("sqlite3", "./highlights.db")
	defer db.Close()
	rows, err := db.Query("SELECT 'uid','text' FROM highlights")
	defer rows.Close()

	highlights := []Highlight{}
	highlight := Highlight{}
	for rows.Next() {
		err = rows.Scan(&highlight.Text, &highlight.Id)
		checkErr(err)

		highlights = append(highlights, highlight)
	}
	return highlights
}

func insertHighlight(db *sql.DB, highlightText string) {
	statement, _ := db.Prepare("INSERT INTO highlights(text) values(?)")
	statement.Exec(highlightText)
}

func createHighlightTable(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS 'highlights' (
		        	'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
			        'text' VARCHAR(144) UNIQUE NOT NULL
				);`)
	_, err = stmt.Exec()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
