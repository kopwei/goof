package goof

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/kopwei/goof/protocols/ofp10"
	"github.com/kopwei/goof/protocols/ofp11"
	"github.com/kopwei/goof/protocols/ofp12"
	"github.com/kopwei/goof/protocols/ofp13"
	"github.com/kopwei/goof/protocols/ofp14"
	"github.com/kopwei/goof/protocols/ofp15"
	"github.com/kopwei/goof/protocols/ofpgeneral"
)

const (
	defaultBufferSize = 50
)

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
	ParseMsg(b []byte) (ofpgeneral.OfpMessage, error)
}

// OfpMessageTunnel is the tunnel of messages in one tcp connection
// between the openflow controller and datapath
type OfpMessageTunnel struct {
	conn net.Conn
	pool *ofpBufferPool
	// Openflow Version
	Version   uint8
	Incomming chan ofpgeneral.OfpMessage
	Outgoing  chan ofpgeneral.OfpMessage
	MsgParser MessageParser
	// Channel on which to receive a shutdown command
	Shutdown chan bool
}

// NewOfpMsgTunnel return the message stream
func NewOfpMsgTunnel(con net.Conn) *OfpMessageTunnel {

	msgTunnel := &OfpMessageTunnel{conn: con}
	msgTunnel.Incomming = make(chan ofpgeneral.OfpMessage)
	msgTunnel.Outgoing = make(chan ofpgeneral.OfpMessage)
	msgTunnel.pool = newBufferPool(defaultBufferSize)
	msgTunnel.MsgParser = nil
	go msgTunnel.sendMessage()
	go msgTunnel.receiveMessage()
	for i := 0; i < defaultBufferSize/2; i++ {
		go msgTunnel.parseWorker()
	}
	return msgTunnel
}

// SendFeatureRequest is used to send the feature request message to datapath
func (mt *OfpMessageTunnel) SendFeatureRequest() {
	header := ofpgeneral.NewOfpHeader(mt.Version)
	switch mt.Version {
	case ofp10.Version:
		header.Type = ofp10.OfpTypeFeaturesRequest
	}
	mt.Outgoing <- header
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
	msg := 0
	hdr := 0
	hdrBuf := make([]byte, 4)

	tmp := make([]byte, 2048)
	buf := <-mt.pool.empty
	for {
		n, err := mt.conn.Read(tmp)
		if err != nil {
			// Handle explicitly disconnecting by closing connection
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			//log.Warnln("InboundError", err)
			//m.Error <- err
			//m.Shutdown <- true
			return
		}
		for i := 0; i < n; i++ {
			if hdr < 4 {
				hdrBuf[hdr] = tmp[i]
				buf.WriteByte(tmp[i])
				hdr++
				if hdr >= 4 {
					msg = int(binary.BigEndian.Uint16(hdrBuf[2:])) - 4
				}
				continue
			}
			if msg > 0 {
				buf.WriteByte(tmp[i])
				msg = msg - 1
				if msg == 0 {
					hdr = 0
					mt.pool.full <- buf
					buf = <-mt.pool.empty
				}
				continue
			}
		}
	}
}

func (mt *OfpMessageTunnel) parseWorker() {
	var err error
	for {
		msgBufBytes := <-mt.pool.full
		if mt.MsgParser == nil {
			mt.MsgParser, err = genMsgParser(msgBufBytes.Bytes())
			if err != nil {
				log.Printf("Message parser generation error %s", err.Error())
			}
		}
		msg, err := mt.MsgParser.ParseMsg(msgBufBytes.Bytes())
		if err != nil {
			log.Printf("Message parsing error %s", err.Error())
		}
		mt.Incomming <- msg
		msgBufBytes.Reset()
		mt.pool.empty <- msgBufBytes
	}
}

func genMsgParser(msg []byte) (MessageParser, error) {
	version, err := ofpgeneral.GetMessageVersion(msg)
	if err != nil {
		return nil, err
	}
	var parser MessageParser
	switch version {
	case ofp10.Version:
		parser = &ofp10.OfpMessageParser{}
	case ofp11.Version:
		parser = &ofp11.OfpMessageParser{}
	case ofp12.Version:
		parser = &ofp12.OfpMessageParser{}
	case ofp13.Version:
		parser = &ofp13.OfpMessageParser{}
	case ofp14.Version:
		parser = &ofp14.OfpMessageParser{}
	case ofp15.Version:
		parser = &ofp15.OfpMessageParser{}
	default:
		parser = nil
		err = fmt.Errorf("Unsupported version %d", version)
	}
	return parser, err
}
