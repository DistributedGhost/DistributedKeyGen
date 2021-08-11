package wdkg
import (
    "github.com/alinush/go-mcl"
    havss  "DistributedKeyGen/havss"
    pred "DistributedKeyGen/prediction"
)

type Wdkg struct {
    //id_havss <--> value of counter
    Counters map[uint32]uint32
    // Algorithm 3: line 2,3
    S        map[uint32] []mcl.Fr
    H        map[uint32] []mcl.Fr
    SP       map[uint32] pred.Prediction

}

func Initialization(n uint32){
    for i := uint32(0); i < n; i++ {
        var s mcl.Fr
        s.Random()
        havss.Create(i, s)
    }
}
