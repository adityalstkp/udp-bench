package message

import (
	"math/rand"
	"testing"
)

func randByte(n int) []byte {
    buff := make([]byte, n)
    for i := range buff {
        buff[i] = byte(rand.Int())
    }

    return buff
}

func BenchmarkPool(b *testing.B) {
    b.ReportAllocs()

    mp := NewMessagePool(1024, 1000000)


    for i := 0; i < b.N; i++ {
        dM := randByte(10)

        go mp.Dequeue(func(msg []byte) {})
        go func ()  {
           mp.Get()
           mp.Enqueue(dM)
        }()

    }
}
