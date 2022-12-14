package poly

import (
	"testing"
	"math"
)

type Polynomial []complex128

// *** Polynomial Multiplication ***
type PolyMulTest struct{
	p,q,expected Polynomial
}

var PolyMulTests = []PolyMulTest {
	PolyMulTest{Polynomial{complex(1,0)},Polynomial{complex(1,0)},Polynomial{complex(1,0)}},
	PolyMulTest{Polynomial{complex(5,0),complex(-2,0),complex(3,0)},Polynomial{complex(3,0),complex(-4,0)},Polynomial{complex(15,0),complex(-26,0),complex(17,0),complex(-12,0)}},
	PolyMulTest{Polynomial{complex(3,0),complex(-4,0)},Polynomial{complex(5,0),complex(-2,0),complex(3,0)},Polynomial{complex(15,0),complex(-26,0),complex(17,0),complex(-12,0)}},
	PolyMulTest{Polynomial{},Polynomial{},Polynomial{}},
	PolyMulTest{Polynomial{complex(1,0)},Polynomial{complex(1,0),complex(1,0)},Polynomial{complex(1,0),complex(1,0)}},
}

func TestPolyMul(t *testing.T){
	for _,test := range PolyMulTests{
		output := PolyMul(test.p, test.q); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}

// *** Polynomial scalar multiplication
type PolyScalarMulTest struct{
	p Polynomial
	a complex128
	expected Polynomial
}

var PolyScalarMulTests = []PolyScalarMulTest{
	PolyScalarMulTest{Polynomial{complex(3,0)},complex(2,0),Polynomial{complex(6,0)}},
	PolyScalarMulTest{Polynomial{complex(7,0),complex(0,0),complex(0,0),complex(0,0),complex(0,0),complex(-4,0)},complex(2,0),Polynomial{complex(14,0),complex(0,0),complex(0,0),complex(0,0),complex(0,0),complex(-8,0)}},
	PolyScalarMulTest{Polynomial{},complex(2,0),Polynomial{}},
}

func TestPolyScalarMul(t *testing.T){
	for _,test := range PolyScalarMulTests{
		output := Poly_scalar_mul(test.p, test.a); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}


// *** Polynomial scaling by degree n. Returns x^n . p(x) ***
type PolyScaleTest struct{
	p Polynomial
	n int
	expected Polynomial
}
var PolyScaleTests = []PolyScaleTest{
	PolyScaleTest{Polynomial{complex(3,0),complex(5,0)},4, Polynomial{complex(3,0),complex(5,0),complex(0,0),complex(0,0),complex(0,0),complex(0,0)}},
	PolyScaleTest{Polynomial{complex(2,0),complex(7,0)},1, Polynomial{complex(2,0),complex(7,0),complex(0,0)}},
	PolyScaleTest{Polynomial{complex(2,0),complex(7,0)},0, Polynomial{complex(2,0),complex(7,0)}},
	PolyScaleTest{Polynomial{},2, Polynomial{}},
	PolyScaleTest{Polynomial{complex(3,0),complex(5,0)},-1, Polynomial{complex(3,0)}},
	PolyScaleTest{Polynomial{complex(2,0),complex(6,0),complex(0,0),complex(0,0),complex(-3,0),complex(-2,0),complex(1,0)},-4, Polynomial{complex(2,0),complex(6,0),complex(0,0)}},
}
func TestPolyScale(t *testing.T){
	for _,test := range PolyScaleTests{
		output := Poly_scale(test.p, test.n); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}

// *** Polynomial extend to n. appends 0's before the polynomail. Doesnt change the polynomial
type PolyExtendTest struct{
	p Polynomial
	n int
	expected Polynomial
}

var PolyExtendTests = []PolyExtendTest{
	PolyExtendTest{Polynomial{complex(2,0)},3,Polynomial{complex(0,0),complex(0,0),complex(2,0)}},
	PolyExtendTest{Polynomial{complex(4,0),complex(0,0),complex(5,0),complex(0,0)},5,Polynomial{complex(0,0),complex(4,0),complex(0,0),complex(5,0),complex(0,0)}},
	PolyExtendTest{Polynomial{complex(4,0),complex(0,0),complex(5,0),complex(0,0)},4,Polynomial{complex(4,0),complex(0,0),complex(5,0),complex(0,0)}},
	PolyExtendTest{Polynomial{},3,Polynomial{complex(0,0),complex(0,0),complex(0,0)}},
}

func TestPolyExtend(t *testing.T){
	for _,test := range PolyExtendTests{
		output := Poly_extend(test.p, test.n); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}


// *** Add two Polynomials
type PolyAddTest struct{
	p,q,expected Polynomial
}
var PolyAddTests = []PolyAddTest{
	PolyAddTest{Polynomial{complex(2,0),complex(0,0),complex(0,0),complex(-7,0),complex(15,0)},Polynomial{complex(-3,0),complex(1,0),complex(0,0),complex(8,0),complex(0,0)},Polynomial{complex(-1,0),complex(1,0),complex(0,0),complex(1,0),complex(15,0)}},
	PolyAddTest{Polynomial{},Polynomial{complex(5,0),complex(0,0)},Polynomial{complex(5,0),complex(0,0)}},
	PolyAddTest{Polynomial{complex(5,0),complex(0,0)},Polynomial{},Polynomial{complex(5,0),complex(0,0)}},
	PolyAddTest{Polynomial{complex(-1,0)},Polynomial{complex(3,0),complex(0,0),complex(0,0)},Polynomial{complex(3,0),complex(0,0),complex(-1,0)}},
}

func TestPolyAdd(t *testing.T){
	for _,test := range PolyAddTests{
		output := Poly_add(test.p, test.q); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}

// *** Subtract two Polynomials
type PolySubTest struct{
	p,q,expected Polynomial
}
var PolySubTests = []PolySubTest{
	PolySubTest{Polynomial{complex(2,0),complex(0,0),complex(0,0),complex(-7,0),complex(15,0)},Polynomial{complex(-3,0),complex(1,0),complex(0,0),complex(8,0),complex(0,0)},Polynomial{complex(5,0),complex(-1,0),complex(0,0),complex(-15,0),complex(15,0)}},
	PolySubTest{Polynomial{},Polynomial{complex(5,0),complex(0,0)},Polynomial{complex(-5,0),complex(0,0)}},
	PolySubTest{Polynomial{complex(5,0),complex(0,0)},Polynomial{},Polynomial{complex(5,0),complex(0,0)}},
	PolySubTest{Polynomial{complex(-1,0)},Polynomial{complex(3,0),complex(0,0),complex(0,0)},Polynomial{complex(-3,0),complex(0,0),complex(-1,0)}},
}

func TestPolySub(t *testing.T){
	for _,test := range PolySubTests{
		output := Poly_sub(test.p, test.q); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}


// *** Reverse Poly Test
type RevPolyTest struct {
	p,expected Polynomial
}
var RevPolyTests = []RevPolyTest{
	RevPolyTest{Polynomial{complex(-5,0),complex(6,0),complex(0,0),complex(7,0),complex(0,0),complex(0,0)},Polynomial{complex(0,0),complex(0,0),complex(7,0),complex(0,0),complex(6,0),complex(-5,0)}},
	RevPolyTest{Polynomial{complex(0,0),complex(0,0),complex(7,0),complex(0,0),complex(6,0),complex(-5,0)},Polynomial{complex(-5,0),complex(6,0),complex(0,0),complex(7,0),complex(0,0),complex(0,0)}},
	RevPolyTest{Polynomial{},Polynomial{}},

}

func TestRevPoly(t *testing.T){
	for _,test := range RevPolyTests{
		output := reverse(test.p); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}


// *** Polynomial Reciprocal Test
type PolyReciprocalTest struct{
	p,expected Polynomial
}

var PolyReciprocalTests = []PolyReciprocalTest{
	PolyReciprocalTest{Polynomial{complex(1,0),complex(1,0)},Polynomial{complex(1,0),complex(-1,0)}},
	PolyReciprocalTest{Polynomial{complex(1,0),complex(-11,0),complex(0,0),complex(-22,0)},Polynomial{complex(1,0),complex(11,0),complex(121,0),complex(1353,0)}},
	// PolyReciprocalTest{Polynomial{complex(1,0),complex(-1,0),complex(1,0), complex(2,0), complex(-1,0),complex(-3,0),complex(1,0),complex(4,0)},Polynomial{complex(1,0), complex(1,0),complex(-3,0),complex(-4,0),complex(3,0),complex(15,0),complex(12,0)}},
}

func TestPolyReciprocal(t *testing.T){
	for _,test := range PolyReciprocalTests{
		output := Reciprocal(test.p); 
		if len(output)!=len(test.expected) {
			t.Errorf("Output and expected different length.")
		}
		for i,val := range output {
			if math.Round(real(val)) != math.Round(real(test.expected[i])){
				t.Errorf("Output %f not equal to expected %f", val, test.expected[i])
				break
			}
		}
	}
}