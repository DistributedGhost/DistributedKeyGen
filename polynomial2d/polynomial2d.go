package polynomial2d

import (
    "fmt"
    poly  "DistributedKeyGen/polynomial"
  mcl "github.com/alinush/go-mcl"
)

type Polynomial2d struct {
    Coeff []mcl.Fr

    // It is degree from 0 to t of recovery polynomial
    DegreeX uint64

    // It is degree from 0 to f of share polynomial
    DegreeY uint64
}

func Create() *Polynomial2d {
    var p2d Polynomial2d
    fmt.Println("DUMMY: polynomial2d Init()")
    return &p2d
}
func Evaluate(x uint64, y uint64) (mcl.Fr, error) {
    var res mcl.Fr
    return res, nil
}
func EvaluateX(x uint64) (poly.Polynomial, error) {
    var coeff poly.Polynomial
    return coeff, nil
}
func EvaluateY(y uint64) (poly.Polynomial, error) {
    var coeff poly.Polynomial
    return coeff, nil
}
