package main

import (
	"github.com/kopwei/goof"
)

func main() {
	ctrler, _ := goof.NewOfpController()
	ctrler.StartListen(6633)
}
