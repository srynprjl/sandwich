package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/srynprjl/sandwich/internal/category"
)

func test() {
	id := 1
	cat := category.Category{
		Id: &id,
	}

	fmt.Print(cat.DoesExists())
}
