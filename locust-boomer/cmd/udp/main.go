package main

import (
	"time"

	"github.com/myzhan/boomer"
)

func udpClientTestCase() {
    start := time.Now()
    time.Sleep(100 * time.Millisecond)
    elapsed := time.Since(start)

    for i := 0; i <= 10; i++ {
        if i % 2 == 0 {
            boomer.RecordSuccess("udp", "bench", elapsed.Nanoseconds()/int64(time.Millisecond), 10)
        } else {
            boomer.RecordFailure("udp", "bench", 10, "example error")
        }
    }
}

func main() {
    udpClientTask := &boomer.Task{
        Name: "udp_client",
        Weight: 10,
        Fn: udpClientTestCase,
    }


    boomer.Run(udpClientTask)
}
