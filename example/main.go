package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tam7t/elfstore"
)

func main() {
	data, err := elfstore.Load()
	if err != nil {
		log.Fatal(err)
	}
	print(data)
	fmt.Println(elfstore.MaxSize())

	// modify
	data[time.Now().String()] = "good times"

	// show result
	fmt.Println("---- after ----")
	print(data)

	// safe changes
	if err := elfstore.Save(data); err != nil {
		log.Fatal(err)
	}
}

func print(data map[string]string) {
	for k, v := range data {
		fmt.Printf("%s:\t %s\n", k, v)
	}
}
