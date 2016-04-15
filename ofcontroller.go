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

// OfpController represents the openflow controller structure
type OfpController interface {
	StartListen(portNo int)
}

type ofpControllerImpl struct {
	bridges []OpenflowSwitch
}

// NewOfpController creates a new openflow controller
func NewOfpController() (OfpController, error) {
	ctrler := &ofpControllerImpl{}
	ctrler.bridges = make([]OpenflowSwitch, 0)
	return ctrler, nil
}

// StartListen will start a tcp listener
func (oc *ofpControllerImpl) StartListen(portNo int) {
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
	msgStream := NewOfpMsgTunnel(conn)
	for {
		select {
		case msg := <-msgStream.Incomming:
			switch m := msg.(type) {

			}
		}

	}
}
