package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
    _ "net/http/pprof"
    
	"github.com/adityalstkp/udp-bench/internal/handler"
	"github.com/adityalstkp/udp-bench/internal/server"
)

var enablePprof bool
var numOfWorkers int

func init() {
    flag.BoolVar(&enablePprof, "enable-pprof", false, "Enable pprof on 6060")
    flag.IntVar(&numOfWorkers, "workers", 4, "Num of message dequeue workers")
}

func main() {
    flag.Parse()

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

    go func ()  {
        for range sig {
            os.Exit(0)
        }
    }()

    if enablePprof {
        go func() {
            log.Println(http.ListenAndServe("localhost:6060", nil))
        }()
    }

    addr := "0.0.0.0:3000"
    udp := server.NewUDPServer(numOfWorkers)
    udp.SetHandler(handler.MessageHandler)
    udp.Listen(addr)
    err := udp.Start(); if err != nil {
        panic(err)
    }

}

