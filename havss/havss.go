package havss

import (
    "net"
    "fmt"
    "github.com/alinush/go-mcl"
    pr "github.com/golang/protobuf/proto"
    c  "DistributedKeyGen/commitment"
    m  "DistributedKeyGen/message"
    poly  "DistributedKeyGen/polynomial"
    poly2d  "DistributedKeyGen/polynomial2d"
    pb_test  "DistributedKeyGen/proto"
)
type MessageType uint64

const (
    Share_t MessageType = iota + 1 // 1
    send_t                         // 2
    echo_t
    ready_t
    shared_t
    reconstruct_t
    reconstruct_share_t
    //out = iota // ?
)


type signature struct {

}
type Havss struct {
    // This variable init by false and changes only
    // Algorithm 1: line 27
    success bool
    id_dealer uint64

    // for each commitment calculates hash by it (uint64)
    // this hash of commitment is key ito map below
    // for every hash exist own counter echo and ready messages
    // these counters of echo and ready are value into map
    echo_counters map[uint64]uint64
    ready_counters map[uint64]uint64
    // A_C from Algorithm 1: line 5
    recovery_points map[uint64]pair
    // B_C from Algorithm 1: line 5
    share_points map[uint64]pair
    // Sig_C from Algorithm 1: line 5
    signatures map[uint64]uint64
}

func send(msg []byte, PORT string) error {
        var CONNECT = "127.0.0.1:" + PORT 
        s, err := net.ResolveUDPAddr("udp4", CONNECT)
        c, err := net.DialUDP("udp4", nil, s)
        if err != nil {
            fmt.Println(err)
            return nil
        }
        fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())

        fmt.Println([]byte(msg))
        _, err = c.Write([]byte(msg))
        defer c.Close()
        return nil
    }
type pair struct {
    Point mcl.Fr
    Sender_index uint64
}
func Create ( i uint64, secret mcl.Fr ) *Havss {
    var h Havss
    h.success = false
    h.id_dealer = i
    h.echo_counters = make(map[uint64]uint64)
    h.ready_counters = make(map[uint64]uint64)
    h.recovery_points = make(map[uint64]pair)
    h.share_points = make(map[uint64]pair)
    h.signatures = make(map[uint64]uint64)
    poly.Init()
    fmt.Println("Mark 1")
    return &h
}

func (h *Havss) Handle(message_type MessageType, msg m.Message) error {
    if message_type == Share_t {
        h.Share_type_processing(msg)
    }
    if message_type == send_t {
        send_type_processing(msg)
    }
    if message_type == echo_t {
        echo_type_processing(msg)
    }
    if message_type == ready_t {
        ready_type_processing(msg)
    }
    if message_type == shared_t {
        shared_type_processing(msg)
    }
    if message_type == reconstruct_t {
        reconstruct_type_processing(msg)
    }
    if message_type == reconstruct_share_t {
        reconstruct_share_type_processing(msg)
    }
    return nil
}

func (h *Havss) Share_type_processing(msg m.Message) error {
    var commit = c.Create(*poly2d.Create())
    var rec_poly, err1 = poly2d.EvaluateX(2)
    if err1 != nil {
        fmt.Println("incorrect return from EvaluateX()")
    }
    var share_poly, err2 = poly2d.EvaluateY(2)
    if err2 != nil {
        fmt.Println("incorrect return from EvaluateY()")
    }
    p := pb_test.SendHAVSS{
        TypeMessage: pb_test.SendHAVSS_SEND,
        IdDealer: h.id_dealer,
        Commitment : commit.ToBytes(),
        RecoveryPoly: rec_poly.ToBytes(),
        SharePoly: share_poly.ToBytes(),
    }
    data, err3 := pr.Marshal(&p)
    if err3 != nil{
        fmt.Println("Failed to convert protobuf to bytes")
    }
    const PORT = "8001"
    send(data, PORT)
    return nil
}
func send_type_processing(msg m.Message) error {
    return nil
}

func echo_type_processing(msg m.Message) error {
    return nil
}

func ready_type_processing(msg m.Message) error {
    return nil
}

func shared_type_processing(msg m.Message) error {
    return nil
}

func reconstruct_type_processing(msg m.Message) error {
    return nil
}

func reconstruct_share_type_processing(msg m.Message) error {
    return nil
}

