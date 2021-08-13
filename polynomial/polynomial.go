package polynomial

import (
  "fmt"
  "encoding/binary"
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

func (p Polynomial) ToBytes() []byte {
    store_degree := make([]byte, 8)
    binary.LittleEndian.PutUint64(store_degree, uint64(p.Degree))
    var store_coeffs []byte
    for i := uint64(0); i < p.Degree; i++ {
        store_coeffs = append(store_coeffs, p.Coeff[i].Serialize()...)
    }
    store_coeffs = append(store_coeffs, store_degree...)
    return store_coeffs
}
