package server

import (
	"errors"
	"net"

	"github.com/adityalstkp/udp-bench/internal/message"
)

type UDPServer struct {
    Address string
    Workers int
    Datagram int
    MaxQueue int
    Handler func(m []byte)
}

func (u UDPServer) Start() error {
    if u.Handler == nil {
        return errors.New("plese provide message handler, handle cannot be nil")
    }

    n, err := net.ListenPacket("udp", u.Address)
    if err != nil {
        return err
    }

    m := message.NewMessage(u.Datagram, u.MaxQueue)

    for i := 0; i < u.Workers; i++ {
        go m.Dequeue(u.Handler)
        go receiveMessage(n, m)
    }
    
    return nil
}

func receiveMessage(c net.PacketConn, m message.Message) {
    defer c.Close()
    
    for {
        msg := m.Pool.Get().([]byte)
        _, _, err := c.ReadFrom(msg)
        if err != nil {
            println(err.Error())
        }
        
        m.Enqueue(msg)
    }
}
