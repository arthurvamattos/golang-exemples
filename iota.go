package main

import "fmt"

func main() {

	const (
		a = iota
		b = iota
		c = iota
	)

	const (
		d = iota * 10
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}