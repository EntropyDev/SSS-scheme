package fft

import (
	"fmt"
	"math"
	"math/cmplx"
)

type Poly []complex128


func FFT(p []complex128, f bool) Poly{
	n := len(p)
	if n == 1 {
		return p
	}
	w_n := cmplx.Exp(complex(0,-2*math.Pi/float64(n)))
	if !f {
		w_n = 1/w_n
	}
	w := complex(1,0)

	pe := make(Poly, n/2)
	po := make(Poly, n/2)

	for i:= 0; i<n; i++ {
		if i%2 == 0 {
			pe[i/2] = p[i]
		}else{
			po[i/2] = p[i]
		}
	}

	ye := FFT(pe, f)
	yo := FFT(po, f)

	y := make(Poly, n)
	if f{
		for j:=0;j<n/2;j++ {
			y[j] = ye[j] + yo[j]*w
			y[j+n/2] = ye[j] - yo[j]*w
			w = w * w_n
		}
	}else{
		for j:=0;j<n/2;j++ {
			y[j] = (ye[j] + yo[j]*w)/2
			y[j+n/2] = (ye[j] - yo[j]*w)/2
			w = w * w_n		
		}	
	}
	return y;
} 

func Hello(name string) string{
	message := fmt.Sprintf("Hi %v, Welcome.",name)
	return message
}