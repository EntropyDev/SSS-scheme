package main

import (
	"fmt"
	// "math"
	// "math/cmplx"
	"fft"
	// "github.com/mjibson/go-dsp/fft"
)

type Poly []complex128

func main(){
	// message := fft.Hello("Vaibhav")
	// fmt.Println(message)
	var p Poly
	p = append(p,complex(8,0))
	p = append(p,complex(-1,1))
	p = append(p,complex(-2,0))
	p = append(p,complex(-1,-1))
	printSlice(p)
	// fmt.Println(math.Sin(2*math.Pi))
	// j := float64(0)
	// n:=4
	// fmt.Println(complex(math.Cos((2*j*math.Pi)/float64(n)),math.Sin((2*j*math.Pi)/float64(n))))

	fmt.Printf("%v",fft.FFT(p,false))
	// fmt.Println(fft.FFT(p))
}

func printSlice(s Poly) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}