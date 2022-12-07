package poly

import (
	// "fmt"
	"fft"
	"math"
	// "math/cmplx"
)

func PolyMul(p []complex128,q []complex128) []complex128{
	degP := len(p)-1
	degQ := len(q)-1
	degC := degP+degQ

	var d int
	d = int(math.Ceil(math.Log2(float64(degC))))
	var z1 = make([]complex128,int(math.Pow(2,float64(d)))-degP+1)
	var z2 = make([]complex128,int(math.Pow(2,float64(d)))-degQ+1)
	p = append(p,z1...)
	q = append(q,z2...)

	pv := fft.FFT(p,true)
	qv := fft.FFT(q,true)
	var cv []complex128 
	for i:=0 ; i< len(pv); i++{
		cv[i] = pv[i]*qv[i]
	}
	c := fft.FFT(cv,false)
	c = c[:degC+1]
	return c
}