package poly

import (
	"fmt"
	"fft"
	"math"
	// "github.com/mjibson/go-dsp/fft"
	// "math/cmplx"
)

// Multiplies two polynomials of degree m,n and returns the resultant polynomial of degree a+b.
// func PolyMul(p []complex128,q []complex128) []complex128{
// 	if len(p) == 0 || len(q) == 0 {
// 		return []complex128{}
// 	}
// 	if len(p) == 1 || len(q) == 1 {
// 		if len(p) == 1 {
// 			return Poly_scalar_mul(q,p[0])
// 		}else{
// 			return Poly_scalar_mul(p,q[0])
// 		}
// 	}
// 	p = reverse(p)
// 	q = reverse(q)
// 	degP := len(p)-1
// 	degQ := len(q)-1
// 	degC := degP+degQ

// 	var d int
// 	d = int(math.Ceil(math.Log2(float64(degC))))
// 	if d>1 {
// 		var z1 = make([]complex128,int(math.Pow(2,float64(d)))-degP-1)
// 	var z2 = make([]complex128,int(math.Pow(2,float64(d)))-degQ-1)
// 	p = append(p,z1...)
// 	q = append(q,z2...)
// 	}
	

// 	pv := fft.FFT(p,true)
// 	qv := fft.FFT(q,true)
// 	var cv []complex128 
// 	for i:=0 ; i< len(pv); i++{
// 		cv = append(cv,pv[i]*qv[i])
// 	}
// 	c := fft.FFT(cv,false)
// 	c = c[:degC+1]
// 	return reverse(c)
// }
func PolyMul(p,q []complex128) []complex128{
	// """Multiply polynomials ``u(x)`` and ``v(x)`` with FFT."""
    if len(p)==0 ||  len(q)==0{
		return make([]complex128,0)
	}
	 p = reverse(p)
	 q = reverse(q)

	d := Poly_deg(p) + Poly_deg(q)
	d = int(math.Ceil(math.Log2(float64(d))))
    U := fft.FFT(Poly_extend(p,d+1),true)
    V := fft.FFT(Poly_extend(q,d+1),true)
	var uv []complex128 
	for i:=0 ; i< len(U); i++{
		uv = append(uv,U[i]*V[i])
	}
	uv = reverse(uv)
    return uv
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
	fmt.Println("Rec Step with k,len(p),cap(p) =",k,len(p),cap(p))
	if len(p)>2{
		q = Reciprocal(p[k/2:])
	}else{
		// p1 := make([]complex128,0)
		// p1 = append(p1,1/p[1])
		q = append(q,1/p[1])
	}
	
	fmt.Println("q(x) = ",q)
	var r []complex128
	q2x := Poly_scalar_mul(q,complex(2,0))
	fmt.Println("-1- ")
	q2x = Poly_scale(q2x,3*k/2-2)
	fmt.Println("-2- ")
	q_square := PolyMul(q,q)
	fmt.Println("-3- ")
	fmt.Println(q_square,p)
	pm := PolyMul(q_square,p)
	fmt.Println("-4- ")
	r = Poly_sub(q2x , pm)
	fmt.Println("-5- ")
	fmt.Println("Rec step ending with k =",k)
	fmt.Println("-6- ")
	return Poly_scale(r,-(k-2))
}

// Takes two polynomials of degree m,n and returns q(x)-quotient and remainder r(x).
// func PolyDiv(p []complex128, q []complex128) ([]complex128, []complex128) {

// }


// ############### HELPER FUNCTIONS
// ---------------


func reverse(slice []complex128) []complex128{
	rev_slice := []complex128{}
	for i:= range slice {
		rev_slice = append(rev_slice,slice[len(slice)-1-i])
	}
	return rev_slice
}

func Poly_deg(p []complex128) int {
	return len(p)-1
}

func Poly_scale(p []complex128, n int) []complex128{
	if len(p) == 0 {
		return []complex128{}
	}
	if n >= 0 {
		return append(p,make([]complex128,n)...)
	}else{
		return p[:len(p)+n]
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

func Poly_add (a []complex128, b []complex128) []complex128{
	var d int
	var res []complex128
	if len(a) > len(b) {
		d = len(a)
		b = Poly_extend(b,d)
	}
	if len(a) < len(b){
		d = len(b)
		a = Poly_extend(a,d)
	}
	for i := range a {
		res = append(res,a[i]+b[i])
	}
	return res
}

func Poly_sub (a []complex128, b []complex128) []complex128{
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