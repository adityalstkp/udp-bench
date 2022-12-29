package handler

import "strings"

func MessageHandler(m []byte) {
    s := string(m)
    println(strings.TrimSuffix(s, "\n"))
}
