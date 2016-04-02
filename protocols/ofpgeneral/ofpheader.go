package ofpgeneral

import (
	"encoding/binary"
	"fmt"
)

// OfpHeader is used to describe the header of OpenFlow message
// Each OpenFlow message begins with the OpenFlow header
type OfpHeader struct {
	// Version is the OFP_VERSION
	Version uint8
	// Type describes one of the OFPT_ constants
	Type uint8
	// Length including this ofp_header
	Length uint16
	// Xid is the transaction id associated with this packet.
	// Replies use the same id as was in the request to facilicate pairing
	Xid uint32
}

// NewOfpHeader creates a reference to a ofp header struct
func NewOfpHeader() *OfpHeader {
	return &OfpHeader{}
}

// UnmarshalBinary transforms the byte array into header data
func (header *OfpHeader) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	header.Version = data[0]
	header.Type = data[1]
	header.Length = binary.BigEndian.Uint16(data[2:4])
	header.Xid = binary.BigEndian.Uint32(data[4:8])
	return nil
}

// MarshalBinary converts the header fields into byte array
func (header *OfpHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8)
	err = nil
	data[0] = header.Version
	data[1] = header.Type
	binary.BigEndian.PutUint16(data[2:4], header.Length)
	binary.BigEndian.PutUint32(data[4:8], header.Xid)
	return data, err
}
