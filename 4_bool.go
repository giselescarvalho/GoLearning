/*go build bool.go
  go run bool.go
*/
package main

import "fmt"

func main() {
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         //false
}
