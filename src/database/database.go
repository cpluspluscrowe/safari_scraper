package database

import (
	"database/sql"
)

type Highlight struct {
	text string
	id   int
}

func InsertHighlights(highlights []string) {
	db := getDatabaseDriver()
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

func GetHighlights(db *sql.DB) []Highlight {
	rows, err := db.Query("SELECT * FROM highlights")
	defer rows.Close()

	highlights := []Highlight{}
	var text string
	var id int
	for rows.Next() {
		err = rows.Scan(&text, &id)
		checkErr(err)

		highlight := Highlight{text, id}
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
			        'text' VARCHAR(144) UNIQUE NOT NULL,
				);`)
	_, err = stmt.Exec()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
