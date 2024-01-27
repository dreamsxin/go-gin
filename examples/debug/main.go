package main

import (
	"log"

	"github.com/dreamsxin/go-gin/debug"
)

type MyData struct {
	IntField   int
	FloatField float64
	StrField   string
	MapField   map[int]string
	SliceField []int
	PointField *MyData
}

func main() {
	debug.Print("abc", 123)

	debug.Pause(true)

	data := &MyData{
		1234,
		77.88,
		"xyz",
		map[int]string{
			1: "abc",
			2: "def",
			3: "ghi",
		},
		[]int{
			3,
			7,
			11,
			13,
			17,
		},
		nil,
	}
	data.PointField = data

	log.Printf("\n%s", debug.Dump(debug.DumpStyle{Pointer: true, Indent: "    "}, data))
	log.Printf("\n%s", debug.Dump(debug.DumpStyle{Format: true, Indent: "    "}, data))

	si := debug.StackTrace(0)
	log.Printf("StackTrace:\n" + si.String("    "))
}
