package main

import (
	"fmt"
	"fft"
)

func main(){
	message := fft.Hello("Vaibhav")
	fmt.Println(message)
}