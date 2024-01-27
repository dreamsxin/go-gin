package main

import (
	"log"

	"github.com/dreamsxin/go-gin/debug"
	"github.com/dreamsxin/go-gin/monitor"
)

func main() {

	log.Printf("\n%s", debug.Dump(debug.DumpStyle{Format: true, Pointer: true, Indent: "    "}, monitor.Monitor()))
}
