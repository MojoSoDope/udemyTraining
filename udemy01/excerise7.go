package main

import "fmt"

func main()  {
	addMe := 0
	for i := 0; i < 1000 ; i++  {

		if i % 3 == 0 || i % 5 == 0 {
			addMe += i
		}
	}

	fmt.Println(addMe)
}
