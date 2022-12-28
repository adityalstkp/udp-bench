package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/adityalstkp/udp-bench/internal/handler"
	"github.com/adityalstkp/udp-bench/internal/server"
)

func main() {
    rCPU := runtime.NumCPU() - 1
    runtime.GOMAXPROCS(rCPU)

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

    uS := server.UDPServer{
        Address: "0.0.0.0:3000",
        Workers: rCPU,
        MaxQueue: 1000000,
        Datagram: 1024,
        Handler: handler.MessageHandler,
    }
    err := uS.Start(); if err != nil {
        panic(err)
    }

    serverCtx, serverStopCtx := context.WithCancel(context.Background())
    go func() {
        <- sig

        serverStopCtx()
	}()

    println("Listening...")

    <- serverCtx.Done()
}

