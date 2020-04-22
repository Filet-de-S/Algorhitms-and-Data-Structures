package main

import (
	"encoding/gob"
	"os"
)

type s struct {
	v int
}

func writeToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func readFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataDecoder := gob.NewDecoder(file)

	var data []int
	err = dataDecoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
