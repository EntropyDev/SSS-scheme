package main

import (
	"fmt"
	// "math"
	// "math/cmplx"
	// "fft"
	"poly"
	// "github.com/mjibson/go-dsp/fft"
)

// type Poly []complex128

func main(){

	var p []complex128
	p = append(p,complex(2,0))
	p = append(p,complex(1,0))
	p = append(p,complex(0,0))
	p = append(p,complex(1,0))
	// printSlice(p)

	var q []complex128
	q = append(q,complex(-4,0))
	q = append(q,complex(0,0))
	q = append(q,complex(1,0))
	// printSlice(q)

	var a []complex128
	a = append(a,complex(7,0))
	a = append(a,complex(-1,0))
	a = append(a,complex(1,0))
	a = append(a,complex(2,0))
	a = append(a,complex(-1,0))
	a = append(a,complex(-3,0))
	a = append(a,complex(1,0))
	a = append(a,complex(4,0))

	fmt.Println(poly.Reciprocal(a))

	// fmt.Println(poly.PolyMul(p,q))

	// fmt.Printf("%v",fft.FFT(p,true))
	// p2 := fft.FFT(p,true)
	// q2 := fft.FFT(q,true)
	// var c2 Poly
	// for i:=0 ; i< len(p2); i++ {
	// 	c2 = append(c2,p2[i]*q2[i])
	// }
	// // var ans []complex128
	// fmt.Println(fft.FFT(c2,false))
	// fmt.Println(fft.IFFT(p))

}

func printSlice(s []complex128) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}