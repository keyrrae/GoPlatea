package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(fibonacci(int64(i)))
	} else {
		fmt.Println(fibonacci(34))
	}
}

func fibonacci(i int64) int64 {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}
