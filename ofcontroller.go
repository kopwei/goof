package goof

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// OfpPacketInMsg is the interface for the openflow package
type OfpPacketInMsg interface{}

// OFApplication defines the openflow application interface
type OFApplication interface {
	// A Switch connected to the controller
	Connected(sw *OpenflowSwitch)

	// Switch disconnected from the controller
	Disconnected(sw *OpenflowSwitch)

	// Controller received a message packet from the switch
	PacketRcvd(sw *OpenflowSwitch, msg OfpPacketInMsg)
}

// StartController will start a tcp listener
func StartController(portNo int) {
	portNoStr := fmt.Sprintf("%d", portNo)
	addr, _ := net.ResolveTCPAddr("tcp", portNoStr)

	var err error
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Println("Listening for connections on", addr)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

}
