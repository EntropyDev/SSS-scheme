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
	// PolyMulTest{[],[],[]},
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

