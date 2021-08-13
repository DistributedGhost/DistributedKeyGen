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

func (commit Commitment) ToBytes() []byte {
// На основе полученного commit создается новый объект
// который хранит байты и возврашается из функции
// объект commit не меняется
    var store []byte
    return store
}
