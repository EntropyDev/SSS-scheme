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
	w := cmplx.Exp(complex(0,float64(2*math.Pi)/float64(n)))
	// if !f{
	// 	w = complex(1/float64(n),0)*cmplx.Pow(cmplx.Exp(complex(0,float64(-2*math.Pi)/float64(n))),-1+0i)
	// 	fmt.Println("working")
	// }
	

	pe := make(Poly, n/2)
	po := make(Poly, n/2)

	for i:= 0; i<n; i++ {
		if i%2 == 0 {
			pe[i/2] = p[i]
		}else{
			po[(i-1)/2] = p[i]
		}
	}
	// fmt.Println(w)
	// fmt.Printf("pe= %v, po = %v",pe,po)
	// return p
	// complex(math.Cos((2*j*math.Pi)/n),math.Sin((2*j*math.Pi)/n))

	ye := FFT(pe, f)
	yo := FFT(po, f)

	// fmt.Printf("n = %d, ye = %v, yo = %v",n,ye,yo)

	y := make(Poly, n)
	if f{
		for j:=0;j<n/2;j++ {
			y[j] = ye[j] + yo[j]*cmplx.Pow(w,complex(float64(j),0))
			y[j+n/2] = ye[j] - yo[j]*cmplx.Pow(w,complex(float64(j),0))
		}
	}else{
		for j:=0;j<n/2;j++ {
			y[j] = ye[j] + yo[j]*complex(1/float64(n),0)*cmplx.Pow(cmplx.Exp(complex(0,float64(-2*math.Pi)/float64(n))),-1+0i)
			y[j+n/2] = ye[j] - yo[j]*complex(1/float64(n),0)*cmplx.Pow(cmplx.Exp(complex(0,float64(-2*math.Pi)/float64(n))),-1+0i)
		}
		
	}
		
	
	return y;

} 

func Hello(name string) string{
	message := fmt.Sprintf("Hi %v, Welcome.",name)
	return message
}