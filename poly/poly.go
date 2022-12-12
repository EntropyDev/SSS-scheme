package poly

import (
	// "fmt"
	"fft"
	"math"
	// "math/cmplx"
)

// Multiplies two polynomials of degree m,n and returns the resultant polynomial of degree a+b.
func PolyMul(p []complex128,q []complex128) []complex128{
	p = reverse(p)
	q = reverse(q)
	degP := len(p)-1
	degQ := len(q)-1
	degC := degP+degQ

	var d int
	d = int(math.Ceil(math.Log2(float64(degC))))
	if d>1 {
		var z1 = make([]complex128,int(math.Pow(2,float64(d)))-degP-1)
	var z2 = make([]complex128,int(math.Pow(2,float64(d)))-degQ-1)
	p = append(p,z1...)
	q = append(q,z2...)
	}
	

	pv := fft.FFT(p,true)
	qv := fft.FFT(q,true)
	var cv []complex128 
	for i:=0 ; i< len(pv); i++{
		cv = append(cv,pv[i]*qv[i])
	}
	c := fft.FFT(cv,false)
	c = c[:degC+1]
	return reverse(c)
}

// Takes polynomial of degree k and returns reciprocal i.e x^2k-2/p(x). 
// a_k-1 != 0 and k is a power of 2
func Reciprocal(p []complex128) []complex128{
	k := len(p);
	if k==1 {
		var a []complex128
		return append(a,1/p[0])
	}
	var q []complex128
	q = Reciprocal(p[:k/2])
	// return q
	var r []complex128
	// tmp := make([]complex128,((3*k/2))-2)
	q2x := multiply(q,complex(2,0))
	q2x = Poly_scale(q2x,3*k/2-2)
	q_square := PolyMul(q,q)
	// return q_square


	r = poly_sub(q2x , PolyMul(q_square,p))
	
	return Poly_scale(r,-k+2)
}

// Takes two polynomials of degree m,n and returns q(x)-quotient and remainder r(x).
// func PolyDiv(p []complex128, q []complex128) ([]complex128, []complex128) {

// }


// ###############
// ###############

func multiply(slice []complex128, factor complex128) []complex128 {
    for i := 0; i < len(slice); i++ {
        slice[i] = slice[i] * factor
    }
	return slice
}

func reverse(slice []complex128) []complex128{
	rev_slice := []complex128{}
	for i:= range slice {
		rev_slice = append(rev_slice,slice[len(slice)-1-i])
	}
	return rev_slice
}

// func diff(a []complex128,b []complex128) []complex128{
// 	sa := len(a)
// 	sb := len(b)
// 	var size_diff int
// 	if sa >= sb { 
// 		size_diff = sa-sb
// 		b = append(make([]complex128,size_diff),b...)
// 	}else{ 
// 		size_diff = sb-sa
// 		a = append(make([]complex128,size_diff),a...)
// 	}
// 	var ans []complex128
// 	for i := range a {
// 		ans = append(ans,a[i]-b[i])
// 	}
// 	return ans
// }

func Poly_deg(p []complex128) int {
	return len(p)-1
}

func Poly_scale(p []complex128, n int) []complex128{
	if n >= 0 {
		return append(p,make([]complex128,n)...)
	}else{
		return p[:n]
	}
}

func Poly_scalar_mul(p []complex128, a complex128) []complex128{
	var res []complex128
	for i:= range p {
		res = append(res,a*p[i])
	}
	return res
}

func Poly_extend(p []complex128, d int) []complex128{
	return append(make([]complex128,d-len(p)),p...)
}

func Poly_norm(p []complex128) []complex128{
	for i,a := range p{
		if a != 0 {
			return p[i:]
		}
	}
	return make([]complex128,0)
}

func poly_add (a []complex128, b []complex128) []complex128{
	var d int
	var res []complex128
	if len(a) > len(b) {
		d = len(a)
		b = Poly_extend(b,d)
	}else{
		d = len(b)
		a = Poly_extend(a,d)
	}
	for i := range a {
		res = append(res,a[i]+b[i])
	}
	return res
}

func poly_sub (a []complex128, b []complex128) []complex128{
	var d int
	var res []complex128
	if len(a) > len(b) {
		d = len(a)
		b = Poly_extend(b,d)
	}else{
		d = len(b)
		a = Poly_extend(a,d)
	}
	for i := range a {
		res = append(res,a[i]-b[i])
	}
	return res
}