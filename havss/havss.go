package havss

import (
    "fmt"
    "github.com/alinush/go-mcl"
    c  "DistributedKeyGen/commitment"
    m  "DistributedKeyGen/message"
    poly  "DistributedKeyGen/polynomial"
    poly2d  "DistributedKeyGen/polynomial2d"
)

type havss struct {
    id_dealer uint32
}
const (
    share_t = iota
    send_t = iota
    echo_t = iota
    ready_t = iota
    shared_t = iota
    reconstruct_t = iota
    reconstruct_share_t = iota
    //out = iota // ?
)
func send (
    type_message uint,
    id_dealer uint32,
    commitment c.Commitment,
    rec poly.Polynomial,
    share poly.Polynomial) error {
    return nil
    }
func Create ( i uint32, secret mcl.Fr ) *havss {
    var h havss
    h.id_dealer = i
    poly.Init()
    var commit = c.Create(*poly2d.Create())
    var rec_poly, err1 = poly2d.EvaluateX(2)
    if err1 != nil {
        fmt.Println("incorrect return from EvaluateX()")
    }
    var share_poly, err2 = poly2d.EvaluateY(2)
    if err2 != nil {
        fmt.Println("incorrect return from EvaluateY()")
    }
    send(send_t, h.id_dealer, *commit, rec_poly, share_poly)
    return &h
}

func Handle(message_type uint, msg m.Message) error {
    if message_type == share_t {
        share_type_processing(msg)
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

func share_type_processing(msg m.Message) error {
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

