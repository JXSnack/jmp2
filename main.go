package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var path = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "me.jxsnack.jp2")
var db *sql.DB

func main() {
	ok := ensureExistence()
	if !ok {
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: jp2 [-m|-d] <name> [<path>]")
		return
	}

	loadRoutes()
	defer safeExit()

	operation := os.Args[1]

	switch operation {
	case "-m":
		callMk()
	case "-d":
		callDel()
	default:
		do()
	}
}

func ensureExistence() bool {
	// ensure folder exists
	ok := os.MkdirAll(path, 0755)

	if ok != nil {
		fmt.Println(ok)
	}

	return ok == nil
}

func loadRoutes() {
	deebee, err := sql.Open("sqlite3", filepath.Join(path, "sqlite3.db"))
	if err != nil {
		panic(err)
	}

	_, err = deebee.Exec("CREATE TABLE IF NOT EXISTS Routes(name TEXT PRIMARY KEY, value TEXT NOT NULL)")
	if err != nil {
		panic(err)
	}

	db = deebee
}

func callMk() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: jp2 -m <name> <path>")
		return
	}

	result, err := db.Exec("INSERT OR IGNORE INTO Routes (name, value) VALUES (?, ?)", os.Args[2], strings.Join(os.Args[3:], " "))
	if err != nil {
		panic(err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		fmt.Println("cancelled, route already exists")
	}
}

func callDel() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: jp2 -d <name>")
		return
	}

	result, err := db.Exec("DELETE FROM Routes WHERE name = ?", os.Args[2])
	if err != nil {
		panic(err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		fmt.Println("cancelled, route didn't exist")
	}
}

func do() {

}

func safeExit() {
	if db != nil {
		db.Close()
	}
}
