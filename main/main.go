package main

import (
	"fmt"
	// "math"
	// "math/cmplx"
	"fft"
)

type Poly []complex128

func main(){
	message := fft.Hello("Vaibhav")
	fmt.Println(message)
	var p Poly
	p = append(p,complex(1,0))
	p = append(p,complex(2,0))
	p = append(p,complex(2,0))
	p = append(p,complex(3,0))
	printSlice(p)
	fmt.Printf("%v",fft.FFT(p))
}

func printSlice(s Poly) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}