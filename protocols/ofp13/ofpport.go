package ofp13

import (
	"bytes"
	"fmt"
	"net"

	"github.com/kopwei/goof/protocols/ofpgeneral"
)

// OFP Port Config
// Flags to indicate behavior of the physical port.  These flags are
// used in ofp_phy_port to describe the current configuration.  They are
// used in the ofp_port_mod message to configure the port's behavior.
const (
	OfpPortConfPortDown = 1 << iota // Port is administratively down.
	OfpPortConfNoSTP                // Disable 802.1D spanning tree on port.
	OfpPortConfNoRecv               /* Drop all packets except 802.1D spanning
	   tree packets. */
	OfpPortConfNoRecvSTP  /* Drop received 802.1D STP packets. */
	OfpPortConfNoFlood    /* Do not include this port when flooding. */
	OfpPortConfNoFwd      /* Drop packets forwarded to port. */
	OfpPortConfNoPacketIn /* Do not send packet-in msgs for port. */
)

// OFP Port State
// Current state of the physical port.  These are not configurable from
// the controller.
const (
	OfpPortStateLinkDown = 1 << 0 /* No physical link present. */

	/* The OfpPortStateSTP* bits have no effect on switch operation.  The
	 * controller must adjust OFPPC_NO_RECV, OFPPC_NO_FWD, and
	 * OFPPC_NO_PACKET_IN appropriately to fully implement an 802.1D spanning
	 * tree. */
	OfpPortStateSTPListen  = iota << 8       /* Not learning or relaying frames. */
	OfpPortStateSTPLearn                     /* Learning but not relaying frames. */
	OfpPortStateSTPForward                   /* Learning and relaying frames. */
	OfpPortStateSTPBlock                     /* Not part of spanning tree. */
	OfpPortStateSTPMask    = (iota - 1) << 8 /* Bit mask for OFPPS_STP_* values. */
)

// OfpPortFeatures
// Features of physical ports available in a datapath.
const (
	OfpPortFeature10MbHD    = 1 << iota /* 10 Mb half-duplex rate support. */
	OfpPortFeature10MbFD                /* 10 Mb full-duplex rate support. */
	OfpPortFeature100MbHD               /* 100 Mb half-duplex rate support. */
	OfpPortFeature100MbFD               /* 100 Mb full-duplex rate support. */
	OfpPortFeature1GbHD                 /* 1 Gb half-duplex rate support. */
	OfpPortFeature1GbFD                 /* 1 Gb full-duplex rate support. */
	OfpPortFeature10GbFD                /* 10 Gb full-duplex rate support. */
	OfpPortFeatureCopper                /* Copper medium. */
	OfpPortFeatureFiber                 /* Fiber medium. */
	OfpPortFeatureAutoNeg               /* Auto-negotiation. */
	OfpPortFeaturePause                 /* Pause. */
	OpfPortFeaturePauseAsym             /* Asymmetric pause. */
)

// OFP Port Reason
// What changed about the physical port
const (
	OfpPortReasonAdd    = iota /* The port was added. */
	OfpPortReasonDelete        /* The port was removed. */
	OfpPortReasonModify        /* Some attribute of the port has changed. */
)

// ofp_error_msg 'code' values for OFPET_PORT_MOD_FAILED.  'data' contains
// at least the first 64 bytes of the failed request. */
// enum ofp_port_mod_failed_code {
const (
	OfpPortModFailedCodeBadPort   = iota /* Specified port does not exist. */
	OfpPortModFailedCodeBadHwAddr        /* Specified hardware address is wrong. */
)

// OfpPhysPort represents the physical port structure
type OfpPhysPort struct {
	PortNo uint16
	HwAddr net.HardwareAddr
	Name   []byte
	config uint32 /* Bitmap of OFPPC_* flags. */
	state  uint32 /* Bitmap of OFPPS_* flags. *

	/* Bitmaps of OpfPortFeature* that describe features.  All bits zeroed if
	 * unsupported or unavailable. */
	Curr       uint32 /* Current features. */
	Advertised uint32 /* Features being advertised by the port. */
	Supported  uint32 /* Features supported by the port. */
	Peer       uint32 /* Features advertised by peer. */
}

// UnmarshalBinary transforms the byte array into body data
func (pp *OfpPhysPort) UnmarshalBinary(data []byte) error {
	if len(data) < 64 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	pp.HwAddr = make([]byte, 6)
	pp.Name = make([]byte, 16)
	return ofpgeneral.UnMarshalFields(buf, &pp.PortNo, &pp.HwAddr, &pp.Name, &pp.config,
		&pp.state, &pp.Curr, &pp.Advertised, &pp.Supported, &pp.Peer)
}

// MarshalBinary converts the header fields into byte array
func (pp *OfpPhysPort) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, pp.PortNo, pp.HwAddr, pp.Name, pp.config,
		pp.state, pp.Curr, pp.Advertised, pp.Supported, pp.Peer); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpPortModMsg represents the message layout of a port modification message
type OfpPortModMsg struct {
	Header ofpgeneral.OfpHeader
	PortNo uint16
	HwAddr net.HardwareAddr /* The hardware address is not
	   configurable.  This is used to
	   sanity-check the request, so it must
	   be the same as returned in an
	   OfpPhyPort struct. */

	Config uint32 /* Bitmap of OFPPC_* flags. */
	Mask   uint32 /* Bitmap of OFPPC_* flags to be changed. */

	Advertise uint32 /* Bitmap of "ofp_port_features"s.  Zero all
	   bits to prevent any action taking place. */
	Padding [4]byte /* Pad to 64-bits. */
}

// UnmarshalBinary transforms the byte array into header data
func (pmm *OfpPortModMsg) UnmarshalBinary(data []byte) error {
	if len(data) < 28 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &pmm.Header, &pmm.PortNo, pmm.HwAddr,
		&pmm.Config, &pmm.Mask, &pmm.Advertise)
}

// MarshalBinary converts the header fields into byte array
func (pmm *OfpPortModMsg) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, pmm.Header, pmm.PortNo, pmm.HwAddr,
		pmm.Config, pmm.Mask, pmm.Advertise); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpPortStatusMsg represents the port status msg structure
/* A physical port has changed in the datapath */
type OfpPortStatusMsg struct {
	Header  ofpgeneral.OfpHeader
	Reason  uint8   /* One of OFPPR_*. */
	Padding [7]byte /* Align to 64-bits. */
	Desc    OfpPhysPort
}

// UnmarshalBinary transforms the byte array into header data
func (psm *OfpPortStatusMsg) UnmarshalBinary(data []byte) error {
	if len(data) < 40 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &psm.Header, &psm.Reason, &psm.Padding,
		&psm.Desc)
}

// MarshalBinary converts the header fields into byte array
func (psm *OfpPortStatusMsg) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, psm.Header, psm.Reason, psm.Padding, psm.Desc); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
