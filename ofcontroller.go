package goof

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
	// listen on all interfaces
	//ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", portNo))

	// accept connection on port
	//conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		//message, _ := bufio.NewReader(conn)
		//go handleMessage(message)
	}
}

func handleMessage() {}
