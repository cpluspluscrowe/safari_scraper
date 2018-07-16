package database

import (
	"database/sql"
	"database/sql/driver"
	//"github.com/mattn/go-sqlite3"
)

type Highlight struct {
	text     string
	id       int
	isPosted int
}

func GetDatabaseDriver() *sql.DB {
	db, err := sql.Open("sqlite3", "./highlights.db")
	if err != nil {
		panic(err)
	}
	checkErr(err)
	createHighlightTable(db)
	return db
}

func createHighlightTable(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS 'highlights' (
		        	'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
			        'text' VARCHAR(144) UNIQUE NOT NULL,
				);`)
	_, err = stmt.Exec(nil)
	checkErr(err)
}

func GetUnpostedHighlights(db *sql.DB) []Highlight {
	rows, err := db.Query("SELECT * FROM highlights where posted = 0")
	defer rows.Close()

	highlights := []Highlight{}
	var text string
	var id int
	var isPosted int
	for rows.Next() {
		err = rows.Scan(&text, &id, &isPosted)
		checkErr(err)

		highlight := Highlight{text, id, isPosted}
		highlights = append(highlights, highlight)
	}
	return highlights
}

func InsertHighlight(db *sql.DB, highlightText []driver.Value) {
	statement, _ := db.Prepare("INSERT INTO highlights(text) values(?)")
	statement.Exec(highlightText)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
