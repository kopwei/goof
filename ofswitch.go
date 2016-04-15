package goof

import (
	"encoding/binary"
	"net"

	"github.com/kopwei/goof/protocols/ofp10"
)

// DatapathID represents the datapath object
type DatapathID struct {
	rawValue uint64
}

// GetRawValue is used to retrieve the raw value of the datapth
func (dpid *DatapathID) GetRawValue() uint64 {
	return dpid.rawValue
}

// GetHwAddr is used to retrieve the mac addr of the datapath
func (dpid *DatapathID) GetHwAddr() net.HardwareAddr {
	hardwareAddr := make([]byte, 8)
	binary.BigEndian.PutUint64(hardwareAddr, dpid.rawValue)
	return hardwareAddr
}

// OpenflowSwitch descibes the switch supports openflow
type OpenflowSwitch interface {
	GetDatapathID() *DatapathID
	DoesSupportOFVer(ofpversion uint8) bool
}

// NewSwitch generates a new switch object
func NewSwitch(tunnel *OfpMessageTunnel, msg *ofp10.OfpSwitchFeatureMsg) OpenflowSwitch {
	// TODO
	return nil
}
