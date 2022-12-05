package fft

import (
	"fmt"
	"math"
	"math/cmplx"
)

type Poly []complex128

func FFT(p []complex128) Poly{
	n := len(p)
	if n == 1 {
		return p
	}
	// w := cmplx.Exp(complex(0,float64(2*math.Pi)/float64(n)))

	pe := make(Poly, n/2)
	po := make(Poly, n/2)

	for i:= 0; i<n; i++ {
		i2 := 0
		if i%2 == 0 {
			pe[i2] = p[i]
		}else{
			po[i2] = p[i]
		}
		i2++
	}

	ye := FFT(pe)
	yo := FFT(po)

	y := make(Poly, n)
	for j:=0;j<n/2;j++ {
		y[j] = ye[j] + yo[j]*cmplx.Exp(complex(0,float64(2*float64(j)*math.Pi)/float64(n)))
		y[j+n/2] = ye[j] - yo[j]*cmplx.Exp(complex(0,float64(2*float64(j)*math.Pi)/float64(n)))
	}
	return y;

} 

func Hello(name string) string{
	message := fmt.Sprintf("Hi %v, Welcome.",name)
	return message
}