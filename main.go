// Package HelloWorld writes a traditional greeting to standard output.
//
// This is an example program to test CI pipelines.
//
package main

import "fmt"

func HelloWorld(who string) string {
	if who == "" {
		who = "World"
	}
	return "Hello " + who + "!"
}

func main() {
	fmt.Println(HelloWorld(""))
}
