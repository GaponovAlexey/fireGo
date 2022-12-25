package main

import (
	"fmt"
	"unsafe"

)

func main() {
  var a uint8 =  22
  var b int =  22
  var c int8 =  22

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))

}
