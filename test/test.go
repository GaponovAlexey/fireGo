package main

import (
	"log"

)

func main() {
	s := "lll"

	log.Println(s)

	newS := &s

	change(newS)

	log.Println(s)

}

func change(str *string) {
	*str = "LOL"
}
