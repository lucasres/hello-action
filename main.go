package main

import "fmt"
import "flag"

func main() {
	fmt.Println("Hello My Little Friend")
}

func init() {
	test := ""
	flag.StringVar(&test, "t", "", "test")
	flag.Parse()
	
	fmt.Println(test)
}
