package message

import "sync"

type Message struct {
    Pool *sync.Pool
    Queue chan []byte
}

func NewMessage(datagram int, maxQueue int) Message {
    return Message{
        Pool: &sync.Pool{
            New: func() interface{} { return make([]byte, datagram) },
        },
        Queue: make(chan []byte, maxQueue),
    }
}

func (m Message) Enqueue(msg []byte) {
    m.Queue <- msg
}

func (m Message) Dequeue(h func(msg []byte)) {
    for msg := range m.Queue {
        h(msg)
        m.Pool.Put(msg)
    }
}
