package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

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

}

func callDel() {

}

func do() {

}
