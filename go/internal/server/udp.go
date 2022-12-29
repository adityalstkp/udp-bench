package server

import (
	"errors"
	"net"

	"github.com/adityalstkp/udp-bench/internal/message"
)

type UDPServer struct {
    Address string
    Workers int
    Handler func(m []byte)
    MessagePool message.IMessagePool
}

func (u UDPServer) Start() error {
    if u.Handler == nil {
        return errors.New("plese provide message handler, handler cannot be nil")
    }

    n, err := net.ListenPacket("udp", u.Address)
    if err != nil {
        return err
    }

    for i := 0; i < u.Workers; i++ {
        go u.MessagePool.Dequeue(u.Handler)
        go u.receiveMessage(n)
    }
    
    return nil
}

func (u UDPServer) receiveMessage(c net.PacketConn) {
    defer c.Close()
    
    for {
        msg := u.MessagePool.Get()
        _, _, err := c.ReadFrom(msg)
        if err != nil {
            println(err.Error())
        }
        
        u.MessagePool.Enqueue(msg)
    }
}
