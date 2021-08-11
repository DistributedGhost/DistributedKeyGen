package polynomial

import (
  "fmt"
  "github.com/alinush/go-mcl"
)

type Polynomial struct {
    Coeff []mcl.Fr
    Degree uint64
}
func CreatePolynomial(degree uint64) *Polynomial {
    var p Polynomial
    p.Degree = degree
    p.Coeff = make([]mcl.Fr, p.Degree)
    for i := uint64(0); i < p.Degree; i++ {
        p.Coeff[i].Random()
    }
    return &p
}
func (r *Polynomial) GetDegree() (uint64, error) {
    return r.Degree, nil
}

func Evaluate(p Polynomial, x mcl.Fr) (mcl.Fr, error) {
    var res mcl.Fr
    err := mcl.FrEvaluatePolynomial(&res, p.Coeff, &x)
    return res, err
}


func Init() {
    var p Polynomial
    p = *CreatePolynomial(5)
    var d, err = p.GetDegree()
    if err == nil {
        fmt.Println("Degree: ", d)
    }
    var x mcl.Fr
    x.SetInt64(5)
    var v, err1 = Evaluate(p, x)
    if err1 == nil {
        fmt.Println("Value p(...) = ", v)
    }
}
