package main

import "fmt"

func main() {
	for i:= 1; i <= 10; i++ {
		for a:=1; a<=i;a++{
			fmt.Printf(" %d", a)
		}
        fmt.Println("")
	}
}

