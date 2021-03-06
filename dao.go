package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main func start!\n")

	var insert_begin int = 1
	var insert_end int = 10000000
	var update_begin int = 9999900
	var update_end int = 10000000
	if update_begin < insert_begin {
		update_begin = insert_begin
	}

	if insert_end < update_end {
		update_end = insert_end
	}

	Init()
	Insert(insert_begin, insert_end)
	//Query()
	Update(update_begin, update_end)
	//Delete()
	Fini()
	fmt.Println("Main func end!\n")
}
