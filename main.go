package main

import (
	"encoding/json"
	"log"
)

func main() {
	data := ReadInputFile()
	p := []Person{}
	if err := json.Unmarshal(data, &p); err != nil {
		log.Panic(err.Error())
	}

	log.Println(p)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
