package ofpgeneral

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
