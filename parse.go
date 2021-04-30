package main

import (
	"io/ioutil"
	"log"
)

func ReadInputFile() (data []byte) {
	var err error
	filename := "input.json"
	if data, err = ioutil.ReadFile(filename); err != nil {
		log.Panicf("Error reading input.json file: %s", err.Error())
	}
	return
}
