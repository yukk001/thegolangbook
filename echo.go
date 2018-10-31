package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t0 :=time.Now().Nanosecond()
	fmt.Println(t0)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	t1 :=time.Now()


	fmt.Println(t1)
}
