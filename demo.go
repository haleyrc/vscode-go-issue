package demo

import (
	"log"
	"os"
)

var w = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds|log.Llongfile)

func Print(s string) {
	w.Println(s)
}
