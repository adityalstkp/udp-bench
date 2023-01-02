package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/myzhan/boomer"
)

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var portFlag string
var targetAddr string

func init() {
    flag.StringVar(&portFlag, "target-port", "3000", "Port bind for BE UDP")
    flag.Parse()

    targetAddr = fmt.Sprintf("0.0.0.0:%s", portFlag)
    println("Target is", targetAddr)
}

func sendPacket() {
    time.Sleep(100 * time.Millisecond)

    start := time.Now()
    
    conn, err := net.Dial("udp", targetAddr)
    if err != nil {
        elapsed := time.Since(start)
        boomer.RecordFailure("udp", "write", elapsed.Milliseconds(), err.Error())
        return 
    }
    defer conn.Close()

    _, err = conn.Write([]byte(randBytesString(10)))
    if err != nil {
        elapsed := time.Since(start)         
        boomer.RecordFailure("udp", "write", elapsed.Milliseconds(), err.Error())
        return
    }

    elapsed := time.Since(start)
    boomer.RecordSuccess("udp", "write", elapsed.Milliseconds(), 0)
    return 
}

func randBytesString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = alphanum[rand.Intn(len(alphanum))]
    }
    return string(b)
}

func main() {
    udpClientTask := &boomer.Task{
        Name: "send_packet",
        Weight: 10,
        Fn: sendPacket,
    }


    boomer.Run(udpClientTask)
}
