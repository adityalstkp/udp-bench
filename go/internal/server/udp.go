package server
    

import (
	"errors"
	"net"
	"sync"

)

const (
    maxQueueSize = 1000000
    UDPPacketSize = 1024
    maxPacketSize = 4096
    packetBufSize = 6 * 1024
)

type UDPServer struct {
    workers int
    handler func(m []byte)
    conn net.PacketConn
    datagramChannel chan []byte
    wait sync.WaitGroup
}

func NewUDPServer(w int) *UDPServer {
    return &UDPServer{
        workers: w,
    } 
}

func (u *UDPServer) SetHandler(h func(m []byte)) {
    u.handler = h
}

func (u *UDPServer) Listen(addr string) error {
    c, err := net.ListenPacket("udp", addr)
    if err != nil {
        return err
    }

    u.conn = c

    return nil
}

func (u *UDPServer) Start() error {
    if u.handler == nil {
        return errors.New("please set a valid handler")
    }

    workers := u.workers
    u.datagramChannel = make(chan []byte, workers)

    u.wait.Add(workers)
    for i := 0; i < workers; i++ {
        go u.parseMessage()
    }

    u.wait.Add(1)
    go u.receiveMessage(u.conn)

    u.wait.Wait()

    return nil
}


func (u *UDPServer) receiveMessage(c net.PacketConn) {
    defer c.Close()
    
    defer u.wait.Done()

    var buf []byte
    
    for {
        if len(buf) < UDPPacketSize {
            buf = make([]byte, packetBufSize, packetBufSize)
        }

        // msg := u.pool.Get().([]byte)
        n, _, err := c.ReadFrom(buf)
        if err != nil {
            println(err.Error())
            continue
        }

        msg := buf[:n]
        buf = buf[n:]
        println(len(buf), cap(buf))
        u.datagramChannel <- msg
    }
}

func (u *UDPServer) parseMessage() {
    defer u.wait.Done()

    for m := range u.datagramChannel {
        u.handler(m)
        // u.pool.Put(m[:maxPacketSize])
    }
}

