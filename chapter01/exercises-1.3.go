package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	num := 10000000

	result1 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < num; i++ {
			echo1(os.Args)
		}
	})
	fmt.Println(result1)

	result2 := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < num; i++ {
			echo2(os.Args)
		}
	})
	fmt.Println(result2)
}

func echo1(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = "\n"
	}
	// fmt.Println(s)
}

func echo2(args []string) {
	strings.Join(args, "\n")
	// fmt.Println(strings.Join(args, "\n"))
}
