package main

import "fmt"

func main()  {
	evenNumber := 0
	for i :=0;  i <= 100; i++ {
		evenNumber = i
		if evenNumber % 2 == 0{
			fmt.Println(evenNumber)

		}
	}
}
