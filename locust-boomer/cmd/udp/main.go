package main

import "github.com/myzhan/boomer"

func udpClientTestCase() {
    boomer.RecordFailure("udp", "example", 10, "example error")
}

func main() {
    udpClientTask := &boomer.Task{
        Name: "udp_client",
        Weight: 10,
        Fn: udpClientTestCase,
    }


    boomer.Run(udpClientTask)
}
