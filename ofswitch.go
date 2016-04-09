package goof

import (
	"encoding/binary"
	"net"
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
