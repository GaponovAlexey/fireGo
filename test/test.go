package main

import "log"

func main() {
	var x = []interface{}{1, 2, 3}
	s := 0
	for s >= 100 {
		s++
		log.Println(x)
		i := 2
		x = append(x, i)
	}
	log.Println(x)

}
