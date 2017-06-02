package main

import "fmt"

func main() {
	var total int = 0
	for i := 0; i < 999; i++ {
		for j := 1; j<999; j++ {
			total +=  (i / j)
		}
	}

	fmt.Println(total)
}

