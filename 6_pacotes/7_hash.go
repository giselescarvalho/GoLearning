// hash NAO CRIP: adler32, crc32, crc64
package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	//criar hash:
	h := crc32.NewIEEE() //1813800674
	//escrever dados:
	h.Write([]byte("codigo com pacote hash ########"))
	//calcular hash
	v := h.Sum32()
	fmt.Println(v)

	i := sha1.New() //[65 90 180 10 233 183 204 78 102 214 118 156 178 192 129 6 232 41 59 72]
	i.Write([]byte("sha1"))
	u := i.Sum([]byte{})
	fmt.Println(u)
}
