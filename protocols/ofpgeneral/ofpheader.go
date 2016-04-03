package ofpgeneral

import (
	"bytes"
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
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, header)
	return err
}

// MarshalBinary converts the header fields into byte array
func (header *OfpHeader) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, *header)
	return buf.Bytes(), err
}
