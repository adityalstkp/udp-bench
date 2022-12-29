package message

import "sync"

type MessagePool struct {
    pool *sync.Pool
    queue chan []byte
}

type IMessagePool interface {
    Enqueue(msg []byte)
    Dequeue(h func(msg []byte))
    Get() []byte
}

func NewMessagePool(datagram int, maxQueue int) IMessagePool {
    return MessagePool{
        pool: &sync.Pool{
            New: func() interface{} { return make([]byte, datagram) },
        },
        queue: make(chan []byte, maxQueue),
    }
}

func (m MessagePool) Enqueue(msg []byte) {
    m.queue <- msg
}

func (m MessagePool) Dequeue(h func(msg []byte)) {
    for msg := range m.queue {
        h(msg)
        m.pool.Put(msg)
    }
}

func (m MessagePool) Get() []byte {
    return m.pool.Get().([]byte)
}
