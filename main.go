package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jmp2 [-m|-d] <name> [<path>]")
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

func callMk() {

}

func callDel() {

}

func do() {

}
