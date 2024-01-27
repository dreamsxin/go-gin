package main

import (
	"fmt"
	"log"

	"github.com/dreamsxin/go-gin/debug"
	"github.com/dreamsxin/go-gin/monitor"
	"github.com/gin-gonic/gin"
)

func main() {

	log.Printf("\n%s", debug.Dump(debug.DumpStyle{Format: true, Pointer: true, Indent: "    "}, monitor.Monitor()))
	r := gin.Default()
	monitorCfg := &monitor.Cfg{
		Status: true,
		Debug:  true,
		//StatusPrefix: "/status",
		StatusHardware: true,
		//StatusHardwarePrefix: "/hardware",
	}
	err := monitor.Register(r, monitorCfg)
	if err != nil {
		fmt.Printf("monitor register err %v\n", err)
		return
	}
	err = r.Run("localhost:8080")
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
