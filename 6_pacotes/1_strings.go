package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("letra", "a"))        //true
	fmt.Println(strings.Contains("computador", "dor")) //true
	fmt.Println(strings.Contains("letra", "z"))        //false

}
