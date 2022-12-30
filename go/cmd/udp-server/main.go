package main

import (
    "context"
    "os"
    "os/signal"
    "runtime"
    "syscall"

    "github.com/adityalstkp/udp-bench/internal/handler"
    "github.com/adityalstkp/udp-bench/internal/message"
    "github.com/adityalstkp/udp-bench/internal/server"
)

func main() {
    rCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(rCPU)

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

    addr := "0.0.0.0:3000"
    mPool := message.NewMessagePool(1024, 1000000)
    uS := server.UDPServer{
        Address: addr,
        Workers: rCPU,
        Handler: handler.MessageHandler,
        MessagePool: mPool,
    }
    err := uS.Start(); if err != nil {
        panic(err)
    }

    serverCtx, serverStopCtx := context.WithCancel(context.Background())
    go func() {
        <- sig

        serverStopCtx()
    }()

    println("Listening on", addr, "with", rCPU, "workers")

    <- serverCtx.Done()
}

