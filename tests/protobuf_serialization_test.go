package havss

import(
    "testing"
    "fmt"
    pr "github.com/golang/protobuf/proto"
    c  "DistributedKeyGen/commitment"
    m  "DistributedKeyGen/message"
   // poly  "DistributedKeyGen/polynomial"
    poly2d  "DistributedKeyGen/polynomial2d"
    pb_test  "DistributedKeyGen/proto"
)

func TestProtobufSerialization(t *testing.T) {
    var commit = c.Create(*poly2d.Create())
    var rec_poly, err1 = poly2d.EvaluateX(2)
    if err1 != nil {
        fmt.Println("incorrect return from EvaluateX()")
    }
    var share_poly, err2 = poly2d.EvaluateY(2)
    if err2 != nil {
        fmt.Println("incorrect return from EvaluateY()")
    }
    id_dealer = 1
    p := pb_test.SendHAVSS{
        TypeMessage: pb_test.SendHAVSS_SEND,
        IdDealer: id_dealer,
        Commitment : commit.ToBytes(),
        RecoveryPoly: rec_poly.ToBytes(),
        SharePoly: share_poly.ToBytes(),
    }
    got, _ := pr.Marshal(&p)

    // sending and receiving process

    var want pb_test.SendHAVSS

    pr.Unmarshal(data, want)
    if got != want {
        t.Errorf("Marshal\Unmarshal process was incorrect")
    }
}
