package ofp13

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/kopwei/goof/protocols/ofpgeneral"
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

// UnmarshalBinary transforms the byte array into body data
func (ao *OfpActionOutput) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &ao.Type, &ao.Len, &ao.Port, &ao.MaxLen)
}

// MarshalBinary converts the header fields into byte array
func (ao *OfpActionOutput) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, ao.Type, ao.Len, ao.Port, ao.MaxLen); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionVlanVID represents action structure for OFPAT_SET_VLAN_VID.
// The VLAN id is 12 bits, so we can use the entire 16 bits to indicate
// special conditions.  All ones is used to match that no VLAN id was
// set.
type OfpActionVlanVID struct {
	Type    uint16 /* OFPAT_SET_VLAN_VID. */
	Len     uint16 /* Length is 8. */
	VlanVID uint16 /* VLAN id. */
	Padding [2]byte
}

// UnmarshalBinary transforms the byte array into body data
func (avv *OfpActionVlanVID) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &avv.Type, &avv.Len, &avv.VlanVID)
}

// MarshalBinary converts the header fields into byte array
func (avv *OfpActionVlanVID) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, avv.Type, avv.Len, avv.VlanVID, avv.Padding); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionVlanPCP represents action structure for OFPAT_SET_VLAN_PCP.
type OfpActionVlanPCP struct {
	Type    uint16 /* OFPAT_SET_VLAN_PCP. */
	Len     uint16 /* Length is 8. */
	VlanPCP uint8  /* VLAN priority. */
	Padding [3]byte
}

// UnmarshalBinary transforms the byte array into body data
func (avp *OfpActionVlanPCP) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &avp.Type, &avp.Len, &avp.VlanPCP)
}

// MarshalBinary converts the header fields into byte array
func (avp *OfpActionVlanPCP) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, avp.Type, avp.Len, avp.VlanPCP, avp.Padding); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionDLAddt represents action structure for OFPAT_SET_DL_SRC/DST.
type OfpActionDLAddt struct {
	Type    uint16 /* OFPAT_SET_DL_SRC/DST. */
	Len     uint16 /* Length is 16. */
	DLAddr  net.HardwareAddr
	Padding [6]byte
}

// UnmarshalBinary transforms the byte array into body data
func (ada *OfpActionDLAddt) UnmarshalBinary(data []byte) error {
	if len(data) < 16 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	if err := ofpgeneral.UnMarshalFields(buf, &ada.Type, &ada.Len); err != nil {
		return err
	}
	ada.DLAddr = make([]byte, 6)
	copy(ada.DLAddr, data[4:])
	return nil
}

// MarshalBinary converts the header fields into byte array
func (ada *OfpActionDLAddt) MarshalBinary() ([]byte, error) {
	data := make([]byte, 16)
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, ada.Type, ada.Len); err != nil {
		return nil, err
	}
	copy(data, buf.Bytes())
	copy(data[4:], ada.DLAddr)
	return data, nil
}

// OfpActionNWAddt represents Action structure for OFPAT_SET_NW_SRC/DST.
type OfpActionNWAddt struct {
	Type   uint16 /* OFPAT_SET_TW_SRC/DST. */
	Len    uint16 /* Length is 8. */
	NWAddr net.IP /* IP address. */
}

// UnmarshalBinary transforms the byte array into body data
func (ana *OfpActionNWAddt) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	if err := ofpgeneral.UnMarshalFields(buf, &ana.Type, &ana.Len); err != nil {
		return err
	}
	ana.NWAddr = make([]byte, 4)
	copy(ana.NWAddr, data[4:])
	return nil
}

// MarshalBinary converts the header fields into byte array
func (ana *OfpActionNWAddt) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, ana.Type, ana.Len); err != nil {
		return nil, err
	}
	copy(data, buf.Bytes())
	copy(data[4:], ana.NWAddr)
	return data, nil
}

// OfpActionTPPort represents action structure for OFPAT_SET_TP_SRC/DST.
type OfpActionTPPort struct {
	Type    uint16 /* OFPAT_SET_TP_SRC/DST. */
	Len     uint16 /* Length is 8. */
	TPPort  uint16 /* TCP/UDP port. */
	Padding [2]byte
}

// UnmarshalBinary transforms the byte array into body data
func (atp *OfpActionTPPort) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &atp.Type, &atp.Len, &atp.TPPort)
}

// MarshalBinary converts the header fields into byte array
func (atp *OfpActionTPPort) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, atp.Type, atp.Len, atp.TPPort, atp.Padding); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionNWToS represents action structure for OFPAT_SET_NW_TOS.
type OfpActionNWToS struct {
	Type    uint16 /* OFPAT_SET_TW_SRC/DST. */
	Len     uint16 /* Length is 8. */
	NWTos   uint8  /* IP ToS (DSCP field, 6 bits). */
	Padding [3]byte
}

// UnmarshalBinary transforms the byte array into body data
func (ant *OfpActionNWToS) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &ant.Type, &ant.Len, &ant.NWTos)
}

// MarshalBinary converts the header fields into byte array
func (ant *OfpActionNWToS) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, ant.Type, ant.Len, ant.NWTos, ant.Padding); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
	Padding [4]byte
}

// UnmarshalBinary transforms the byte array into header data
func (ah *OfpActionHeader) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &ah.Type, &ah.Len)
}

// MarshalBinary converts the header fields into byte array
func (ah *OfpActionHeader) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, ah.Type, ah.Len, ah.Padding); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionEnqueueInfo represents the OFPAT_ENQUEUE action struct: send packets to given queue on port.
type OfpActionEnqueueInfo struct {
	Type uint16 /* OFPAT_ENQUEUE. */
	Len  uint16 /* Len is 16. */
	Port uint16 /* Port that queue belongs. Should
	   refer to a valid physical port
	   (i.e. < OFPP_MAX) or OFPP_IN_PORT. */
	Padding [6]byte /* Pad for 64-bit alignment. */
	QueueID uint32  /* Where to enqueue the packets. */
}

// UnmarshalBinary transforms the byte array into body data
func (aei *OfpActionEnqueueInfo) UnmarshalBinary(data []byte) error {
	if len(data) < 16 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &aei.Type, &aei.Len, &aei.Port, &aei.Padding, &aei.QueueID)
}

// MarshalBinary converts the header fields into byte array
func (aei *OfpActionEnqueueInfo) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, aei.Type, aei.Len, aei.Port, aei.Padding, aei.QueueID); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpActionMsg represents the body structure of the action msg sent to datapath
type OfpActionMsg struct {
	Header OfpActionHeader
	Body   ofpgeneral.OfpMessage
}

// UnmarshalBinary transforms the byte array into msg data
func (oam *OfpActionMsg) UnmarshalBinary(data []byte) error {
	if err := (&oam.Header).UnmarshalBinary(data); err != nil {
		return err
	}
	dataStartIdx := 4
	actionType := binary.BigEndian.Uint16(data[dataStartIdx : dataStartIdx+2])
	switch actionType {
	case OfpActionOutputToPort:
		oam.Body = &OfpActionOutput{}
	case OfpActionSetVlanVID:
		oam.Body = &OfpActionVlanVID{}
	case OfpActionSetVlanPCP:
		oam.Body = &OfpActionVlanPCP{}
	case OfpActionSetDLSrc:
		oam.Body = &OfpActionDLAddt{}
	case OfpActionSetDLDst:
		oam.Body = &OfpActionDLAddt{}
	case OfpActionSetNWSrc:
		oam.Body = &OfpActionNWAddt{}
	case OfpActionSetNWDst:
		oam.Body = &OfpActionNWAddt{}
	case OfpActionSetNWToS:
		oam.Body = &OfpActionNWToS{}
	case OfpActionSetTPSrc:
		oam.Body = &OfpActionTPPort{}
	case OfpActionSetTPDst:
		oam.Body = &OfpActionTPPort{}
	case OfpActionEnqueue:
		oam.Body = &OfpActionEnqueueInfo{}
	}
	oam.Body.UnmarshalBinary(data[dataStartIdx:])
	return nil
}

// MarshalBinary transforms the msg data into byte array
func (oam *OfpActionMsg) MarshalBinary() ([]byte, error) {
	data := make([]byte, oam.Header.Len)
	headerData, err := oam.Header.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(data, headerData)
	bodyData, err := oam.Body.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(data[4:], bodyData)
	return data, nil
}
