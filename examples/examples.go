package main

import (
	"github.com/kopwei/goof"
)

func main() {
	// Launch a openflow controller instance and start listen
	ctrler, _ := goof.NewOfpController()
	ctrler.StartListen(6633)
}
