// io  :::  https://pkg.go.dev/io
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Hello, Gophers")
	err := ioutil.WriteFile("HelloWorld.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
