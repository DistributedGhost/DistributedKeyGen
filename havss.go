package main

import "fmt"
import "github.com/alinush/go-mcl"


type polynomial struct {
    Coeff []mcl.Fr
    Degree int
}

func (r *polynomial) degree() (int, error) {
    return r.Degree, nil
}
func (r *polynomial) coeff() (error) {
    for i := 0; i < 2; i++ {
        fmt.Println(r.Coeff[i])
    }
    return nil
}

func generateFr(count uint64) []mcl.Fr {
	base := make([]mcl.Fr, count)
	for i := uint64(0); i < count; i++ {
		base[i].Random()
	}
	return base
}

func main(){
    havss_init()
}
func havss_init() {
    c := make([]mcl.Fr, 0, 2)
    copy(c, generateFr(2))
    var p polynomial
    p = polynomial{Coeff: c, Degree: 2}
    fmt.Println(p.degree())
    fmt.Println(p.coeff())
//    println(p.degree())
}
