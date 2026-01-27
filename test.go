package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/srynprjl/sandwich/internal/category"
)

func test() {
	id := 3
	title := "Hello"
	shorthand := "hi"
	cat := category.Category{
		Id:        &id,
		Title:     &title,
		Shorthand: &shorthand,
	}

	fmt.Print(cat.Update())
}
