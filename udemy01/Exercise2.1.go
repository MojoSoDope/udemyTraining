package main

import "fmt"

func chopIt(x int) (int,bool)  {
	return x / 2, x % 2 == 0
}

func main() {
	n, even := chopIt(5)
	fmt.Println(n, even)
}