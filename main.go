package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ok := ensureExistence()
	if !ok {
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: jp2 [-m|-d] <name> [<path>]")
		return
	}

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
	ok := os.MkdirAll(filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "me.jxsnack.jmp2"), 0666)

	if ok != nil {
		fmt.Println(ok)
	}

	return ok == nil
}

func callMk() {

}

func callDel() {

}

func do() {

}
