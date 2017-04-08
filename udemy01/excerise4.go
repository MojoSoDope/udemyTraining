package main

import "fmt"

func main()  {
	var smallNumber int
	var largeNumber int
	fmt.Print("Please enter a larger number")
	fmt.Scan(&largeNumber)
	fmt.Print("Please enter a small number")
	fmt.Scan(&smallNumber)
	var totalAmount = largeNumber % smallNumber
	fmt.Printf("When you divide" + largeNumber + "by " + smallNumber + " you get " + totalAmount)
}

