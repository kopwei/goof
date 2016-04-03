package ofp10

import (
	"encoding/binary"
	"fmt"
	"net"
)

// OFP Action Type
const (
	OfpActionOutputToPort = iota /* Output to switch port. */
	OfpActionSetVlanVID          /* Set the 802.1q VLAN id. */
	OfpActionSetVlanPCP          /* Set the 802.1q priority. */
	OfpActionStripVlan           /* Strip the 802.1q header. */
	OfpActionSetDLSrc            /* Ethernet source address. */
	OfpActionSetDLDst            /* Ethernet destination address. */
	OfpActionSetNWSrc            /* IP source address. */
	OfpActionSetNWDst            /* IP destination address. */
	OfpActionSetNWToS            /* IP ToS (DSCP field, 6 bits). */
	OfpActionSetTPSrc            /* TCP/UDP source port. */
	OfpActionSetTPDst            /* TCP/UDP destination port. */
	OfpActionEnqueue             /* Output to queue.  */
	OfpActionVendor       = 0xffff
)

// ofp_error_msg 'code' values for OFPET_BAD_ACTION.  'data' contains at least
// the first 64 bytes of the failed request. */
//enum ofp_bad_action_code {
const (
	OfpBadActionCodeBadType       = iota /* Unknown action type. */
	OfpBadActionCodeBadLen               /* Length problem in actions. */
	OfpBadActionCodeBadVendor            /* Unknown vendor id specified. */
	OfpBadActionCodeBadVendorType        /* Unknown action type for vendor id. */
	OfpBadActionCodeBadOutPort           /* Problem validating output action. */
	OfpBadActionCodeBadArgument          /* Bad action argument. */
	OfpBadActionCodeErrPerm              /* Permissions error. */
	OfpBadActionCodeTooMany              /* Can't handle this many actions. */
	OfpBadActionCodeBadQueue             /* Problem validating output queue. */
)

// OfpActionOutput represents the ofp action output
// Action structure for OFPAT_OUTPUT, which sends packets out 'port'.
// When the 'port' is the OFPP_CONTROLLER, 'max_len' indicates the max
// number of bytes to send.  A 'max_len' of zero means no bytes of the
// packet should be sent
type OfpActionOutput struct {
	Type   uint16 /* OFPAT_OUTPUT. */
	Len    uint16 /* Length is 8. */
	Port   uint16 /* Output port. */
	MaxLen uint16 /* Max length to send to controller. */
}

// OfpActionVlanVID represents action structure for OFPAT_SET_VLAN_VID.
// The VLAN id is 12 bits, so we can use the entire 16 bits to indicate
// special conditions.  All ones is used to match that no VLAN id was
// set.
type OfpActionVlanVID struct {
	Type    uint16 /* OFPAT_SET_VLAN_VID. */
	Len     uint16 /* Length is 8. */
	VlanVID uint16 /* VLAN id. */
	//uint8_t pad[2];
}

// OfpActionVlanPCP represents action structure for OFPAT_SET_VLAN_PCP.
type OfpActionVlanPCP struct {
	Type    uint16 /* OFPAT_SET_VLAN_PCP. */
	Len     uint16 /* Length is 8. */
	VlanPCP uint8  /* VLAN priority. */
	//uint8_t pad[3];
}

// OfpActionDLAddt represents action structure for OFPAT_SET_DL_SRC/DST.
type OfpActionDLAddt struct {
	Type   uint16 /* OFPAT_SET_DL_SRC/DST. */
	Len    uint16 /* Length is 16. */
	DLAddr net.HardwareAddr
	//uint8_t pad[6];
}

// OfpActionNWAddt represents Action structure for OFPAT_SET_NW_SRC/DST.
type OfpActionNWAddt struct {
	Type   uint16 /* OFPAT_SET_TW_SRC/DST. */
	Len    uint16 /* Length is 8. */
	NWAddr net.IP /* IP address. */
}

// OfpActionTPPort represents action structure for OFPAT_SET_TP_SRC/DST.
type OfpActionTPPort struct {
	Type   uint16 /* OFPAT_SET_TP_SRC/DST. */
	Len    uint16 /* Length is 8. */
	TPPort uint16 /* TCP/UDP port. */
	//uint8_t pad[2];
}

// OfpActionNWToS represents action structure for OFPAT_SET_NW_TOS.
type OfpActionNWToS struct {
	Type  uint16 /* OFPAT_SET_TW_SRC/DST. */
	Len   uint16 /* Length is 8. */
	NWTos uint8  /* IP ToS (DSCP field, 6 bits). */
	//uint8_t pad[3];
}

// OfpActionVendorHeader represents action header for OFPAT_VENDOR.
// The rest of the body is vendor-defined.
type OfpActionVendorHeader struct {
	Type   uint16 /* OFPAT_VENDOR. */
	Len    uint16 /* Length is a multiple of 8. */
	Vendor uint32 /* Vendor ID, which takes the same form
	   as in "struct ofp_vendor_header". */
}

// OfpActionHeader represents the header structure that is common to all actions.
// The length includes the header and any padding used to make the action 64-bit aligned.
// NB: The length of an action *must* always be a multiple of eight.
type OfpActionHeader struct {
	Type uint16 /* One of OFPAT_*. */
	Len  uint16 /* Length of action, including this
	   header.  This is the length of action,
	   including any padding to make it
	   64-bit aligned. */
	//uint8_t pad[4];
}

// UnmarshalBinary transforms the byte array into header data
func (ah *OfpActionHeader) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	ah.Type = binary.BigEndian.Uint16(data[:2])
	ah.Len = binary.BigEndian.Uint16(data[2:4])

	return nil
}

// MarshalBinary converts the header fields into byte array
func (ah *OfpActionHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8)
	err = nil
	data[0] = header.Version
	data[1] = header.Type
	binary.BigEndian.PutUint16(data[2:4], header.Length)
	binary.BigEndian.PutUint32(data[4:8], header.Xid)
	return data, err
}

// OfpActionEnqueueInfo represents the OFPAT_ENQUEUE action struct: send packets to given queue on port.
type OfpActionEnqueueInfo struct {
	Type uint16 /* OFPAT_ENQUEUE. */
	Len  uint16 /* Len is 16. */
	Port uint16 /* Port that queue belongs. Should
	   refer to a valid physical port
	   (i.e. < OFPP_MAX) or OFPP_IN_PORT. */
	//uint8_t pad[6];           /* Pad for 64-bit alignment. */
	QueueID uint32 /* Where to enqueue the packets. */
}
