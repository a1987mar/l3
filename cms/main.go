package main

import (
	"fmt"
	"les3/pkg/documentstore"
)

func main() {
	d1 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d1.Fields["Name1"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "event1.go",
	}

	d2 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d2.Fields["Name2"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "event2.go",
	}

	d3 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d3.Fields["Name3"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeBool,
		Value: true,
	}

	d4 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d4.Fields["Name4"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "event4.go",
	}

	t := documentstore.Document{}
	t.Put(d1)
	t.Put(d2)
	t.Put(d3)
	t.Put(d4)

	doc, res := t.Get("Name1")
	fmt.Println("get", res, doc)

	res = t.Delete("Name4")
	fmt.Println("DELETE", res)

	fmt.Println(t.List())

}
