package commitment

import (
    p "DistributedKeyGen/polynomial2d"
)

type Commitment struct {

}


func Create(coeff p.Polynomial2d) *Commitment {
    var c Commitment
    return &c
}
