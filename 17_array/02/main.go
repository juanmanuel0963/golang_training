package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = fmt.Sprint(i) //string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
