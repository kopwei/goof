package ofpgeneral

import (
	"log"
	"net"
)

// OfpMessage is the general representation of the messages
// transferred between controller and switch
type OfpMessage interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}

// OfpMessageTunnel is the tunnel of messages in one tcp connection
// between the openflow controller and datapath
type OfpMessageTunnel struct {
	conn net.Conn
	// Openflow Version
	Version   uint8
	Incomming chan OfpMessage
	Outgoing  chan OfpMessage
}

// NewOfpMsgTunnel return the message stream
func NewOfpMsgTunnel(con net.Conn) *OfpMessageTunnel {
	msgTunnel := &OfpMessageTunnel{conn: con}
	msgTunnel.Incomming = make(chan OfpMessage)
	msgTunnel.Outgoing = make(chan OfpMessage)
	go msgTunnel.sendMessage()
	go msgTunnel.receiveMessage()
	return msgTunnel
}

func (mt *OfpMessageTunnel) sendMessage() {
	for {
		msg := <-mt.Outgoing
		data, _ := msg.MarshalBinary()
		if _, err := mt.conn.Write(data); err != nil {
			log.Printf("Error in sending messages %s", err.Error())
		}
	}
}

func (mt *OfpMessageTunnel) receiveMessage() {
	msgbuf := make([]byte, 2048)
	for {
		n, err := mt.conn.Read(msgbuf)
		if err != nil {
			log.Printf("Error in receiving the msg, %s", err.Error())
		}
	}
}
