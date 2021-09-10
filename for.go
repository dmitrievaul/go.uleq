package main

import "fmt"

func main() {
	i := 10
	for i <= 6 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 8; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("yo privet")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%10 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
