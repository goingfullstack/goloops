package main

import (
	"fmt"
)

type Element struct {
	Id string
}

var (
	elements []Element = []Element{
		Element{Id: "1"},
		Element{Id: "2"},
		Element{Id: "3"},
	}
)

func main() {
	m := makeMap()

	fmt.Println(m["1"].Id)
	fmt.Println(m["2"].Id)
	fmt.Println(m["3"].Id)
}

func makeMap() map[string]*Element {
	m := make(map[string]*Element)
	for _, e := range elements {
		m[e.Id] = &e
	}
	return m
}
