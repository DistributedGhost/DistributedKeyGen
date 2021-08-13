package wdkg
import (
    "fmt"
    "github.com/alinush/go-mcl"
    m  "DistributedKeyGen/message"
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

func Initialization(n uint64){
    //var havsses []havss.Havss
    //havsses := make([]havss.Havss, 0, n)
    s := new (mcl.Fr)
    s.Random()
    h :=  *havss.Create(2, *s)

        fmt.Println("DEBUG: Mark 2")
        var msg m.Message
        var messageType havss.MessageType
        messageType = havss.Share_t
        h.Handle(messageType, msg)
        fmt.Println("DEBUG: Mark 2")
//        fmt.Println(len(havsses))
//    for i := uint64(0); i < n; i++ {
//        fmt.Println("Mark 2")
//        var s mcl.Fr
//        s.Random()
//        havsses[i] = *havss.Create(i, s)
//    }
}
