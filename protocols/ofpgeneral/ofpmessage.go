package ofpgeneral

import (
	"bytes"
	"io"
	"log"
	"net"
)

const (
	defaultBufferSize = 50
)

// OfpMessage is the general representation of the messages
// transferred between controller and switch
type OfpMessage interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}

// ofpBufferPool is the message buffer pool
type ofpBufferPool struct {
	empty chan *bytes.Buffer
	full  chan *bytes.Buffer
}

// newBufferPool creates the message buffer pool
func newBufferPool(size int) *ofpBufferPool {
	m := &ofpBufferPool{}
	m.empty = make(chan *bytes.Buffer, size)
	m.full = make(chan *bytes.Buffer, size)
	for i := 0; i < size; i++ {
		m.empty <- bytes.NewBuffer(make([]byte, 0, 2048))
	}
	return m
}

// MessageParser is the interface for message parser
type MessageParser interface {
	ParseMsg(b []byte) (OfpMessage, error)
}

// OfpMessageTunnel is the tunnel of messages in one tcp connection
// between the openflow controller and datapath
type OfpMessageTunnel struct {
	conn net.Conn
	pool *ofpBufferPool
	// Openflow Version
	Version   uint8
	Incomming chan OfpMessage
	Outgoing  chan OfpMessage
	MsgParser MessageParser
}

// NewOfpMsgTunnel return the message stream
func NewOfpMsgTunnel(con net.Conn, parser MessageParser) *OfpMessageTunnel {

	msgTunnel := &OfpMessageTunnel{conn: con}
	msgTunnel.Incomming = make(chan OfpMessage)
	msgTunnel.Outgoing = make(chan OfpMessage)
	msgTunnel.pool = newBufferPool(defaultBufferSize)
	msgTunnel.MsgParser = parser
	go msgTunnel.sendMessage()
	go msgTunnel.receiveMessage()
	for i := 0; i < defaultBufferSize/2; i++ {
		go msgTunnel.parseWorker()
	}
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
	msg := bytes.NewBuffer(nil)
	msgbuf := make([]byte, 2048)
	for {
		n, err := mt.conn.Read(msgbuf)
		msg.Write(msgbuf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error in receiving the msg, %s", err.Error())
		}
	}
	msg.Reset()
}

func (mt *OfpMessageTunnel) parseWorker() {
	for {
		msgBufBytes := <-mt.pool.full
		msg, err := mt.MsgParser.ParseMsg(msgBufBytes.Bytes())
		if err != nil {
			log.Printf("Message parsing error %s", err.Error())
		}
		mt.Incomming <- msg
		msgBufBytes.Reset()
		mt.pool.empty <- msgBufBytes
	}
}
