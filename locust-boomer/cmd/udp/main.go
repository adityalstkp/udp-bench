package main

import (
	"flag"
	"fmt"
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
    start := time.Now()
    
    conn, err := net.Dial("udp", targetAddr)
    if err != nil {
        elapsed := time.Since(start)
        boomer.RecordFailure("udp", "write", elapsed.Milliseconds(), err.Error())
        return 
    }
    defer conn.Close()

    _, err = conn.Write([]byte(randBytes(10)))
    if err != nil {
        elapsed := time.Since(start)         
        boomer.RecordFailure("udp", "write", elapsed.Milliseconds(), err.Error())
        return
    }

    elapsed := time.Since(start)
    boomer.RecordSuccess("udp", "write", elapsed.Milliseconds(), 0)
    return 
}

func randBytes(n int) []byte {
    b := make([]byte, n)
    for i, v := range b {
        b[i] = alphanum[v%byte(len(alphanum))]
    }
    return b
}

func main() {
    udpClientTask := &boomer.Task{
        Name: "send_packet",
        Weight: 10,
        Fn: sendPacket,
    }


    boomer.Run(udpClientTask)
}
