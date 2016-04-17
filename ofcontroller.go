package goof

import (
	"fmt"
	"net"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/kopwei/goof/protocols/ofp10"
	"github.com/kopwei/goof/protocols/ofp15"
	"github.com/kopwei/goof/protocols/ofpgeneral"
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
	hello := ofpgeneral.NewHelloMsg(4)
	msgStream.Outgoing <- hello
	for {
		select {
		case msg := <-msgStream.Incomming:
			switch m := msg.(type) {
			case *ofpgeneral.OfpHelloMsg:
				version, _ := ofpgeneral.GetOfpMsgVersion(m)
				if isVersionValid(version) {
					msgStream.Version = version
					msgStream.SendFeatureRequest()
				} else {
					// Connection should be severed if controller
					// doesn't support switch version.
					log.Println("Received unsupported ofp version", version)
					msgStream.Shutdown <- true
				}

			// After a vaild FeaturesReply has been received we
			// have all the information we need. Create a new
			// switch object and notify applications.
			case *ofp10.OfpSwitchFeatureMsg:
				log.Printf("Received ofp1.0 Switch feature response: %+v", *m)

				// Create a new switch and handover the stream
				NewSwitch(msgStream, m)

				// Let switch instance handle all future messages..
				return

			// An error message may indicate a version mismatch. We
			// disconnect if an error occurs this early.
			case *ofpgeneral.OfpErrMsg:
				log.Warnf("Received  error msg: %+v", *m)
				msgStream.Shutdown <- true
			}
		case err := <-msgStream.Error:
			// The connection has been shutdown.
			log.Println(err)
			return
		case <-time.After(time.Second * 3):
			// This shouldn't happen. If it does, both the controller
			// and switch are no longer communicating. The TCPConn is
			// still established though.
			log.Warnln("Connection timed out.")
			return
		}
	}
}

func isVersionValid(v uint8) bool {
	return v > 0 && v < ofp15.Version
}
