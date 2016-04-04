package ofp10

import (
	"encoding/binary"

	"github.com/kopwei/goof/protocols/ofpgeneral"
)

const (
	// Version is the value of version byte in ofp header
	Version = 0x01
)

// OfpConfigFlags
const (
	/* Handling of IP fragments. */
	OfpConfFragNormal = iota /* No special handling for fragments. */
	OfpConfFragDrop          /* Drop fragments. */
	OfpConfFragReAsm         /* Reassemble (only if OFPC_IP_REASM set). */
	OfpConfFragMask
)

// Ofp Capability flags
// Capabilities supported by the datapath.
const (
	OfpCapFlowStats  = 1 << iota /* Flow statistics. */
	OfpCapTableStats             /* Table statistics. */
	OfpCapPortStats              /* Port statistics. */
	OfpCapSTP                    /* 802.1d spanning tree. */
	OfpCapReserved               /* Reserved, must be zero. */
	OfpCapIPReAsm                /* Can reassemble IP fragments. */
	OfpQueueStats                /* Queue statistics. */
	OfpArpMatchIP                /* Match IP addresses in ARP pkts. */
)

// The OFP Type constants
const (
	/* Immutable Messages*/
	OfpTypeHello = iota /* Symmetric message */
	OfpTypeError        /* Symmetric message */
	OfpTypeEchoRequest
	OfpTypeEchoReply
	OfpTypeExperimenter

	/* Swiitch_configuration messages */
	OfpTypeFeaturesRequest
	OfpTypeFeaturesReply
	OfpTypeGetConfigRequest
	OfpTypeGetConfigReply
	OfpTypeSetConfig

	/* Asynchronous message */
	OfpTypePacketIn
	OfpTypeFlowRemoved
	OfpTypePortStatus

	/* Controller command messages */
	OfpTypePacketOut
	OfpTypeFlowMod
	OfpTypeGroupMod
	OfpTypePortMod
	OfpTypeTableMod

	/* Multipart messages */
	OfpTypeMultiPartRequest
	OfpTypeMultiPartReply

	/* Barrier messages */
	OfpTypeBarrierRequest
	OfpTypeBarrierReply

	/* Queue Configuration messages. */
	OfpTypeQueueGetConfigRequest /* Controller/switch message */
	OfpTypeQueueGetConfigReply   /* Controller/switch message */
)

// Values for 'type' in ofp_error_message.  These values are immutable: they
// will not change in future versions of the protocol (although new values may
// be added).
//enum ofp_error_type {
const (
	OfpErrTypeHelloFailed   = iota /* Hello protocol failed. */
	OfpErrTypeBadRequest           /* Request was not understood. */
	OfpErrTypeBadAction            /* Error in action description. */
	OfpErrTypeFlowModFailed        /* Problem modifying flow entry. */
	OfpErrTypePortModFailed        /* Port mod request failed. */
	OfpErrTypeQueueOpFailed        /* Queue operation failed. */
)

// ofp_error_msg 'code' values for OFPET_HELLO_FAILED.  'data' contains an
// ASCII text string that may give failure details. */
//enum ofp_hello_failed_code {
const (
	OfpHelloFaildCodeIncompatioble = iota /* No compatible version. */
	OfpHelloFaildCodeErrPerm              /* Permissions error. */
)

// ofp_error_msg 'code' values for OFPET_BAD_REQUEST.  'data' contains at least
// the first 64 bytes of the failed request.
//enum ofp_bad_request_code {
const (
	OfpBadReqCodeBadVersion    = iota /* ofp_header.version not supported. */
	OfpBadReqCodeBadType              /* ofp_header.type not supported. */
	OfpBadReqCodeBadStat              /* ofp_stats_request.type not supported. */
	OfpBadReqCodeBadVendor            /* Vendor not supported (in ofp_vendor_header or ofp_stats_request or ofp_stats_reply). */
	OfpBadReqCodeBadSubType           /* Vendor subtype not supported. */
	OfpBadReqCodeErrPerm              /* Permissions error. */
	OfpBadReqCodeBadLen               /* Wrong request length for type. */
	OfpBadReqCodeBufferEmpty          /* Specified buffer has already been used. */
	OfpBadReqCodeBufferUnknown        /* Specified buffer does not exist. */
)

// ofp_error msg 'code' values for OFPET_QUEUE_OP_FAILED. 'data' contains
// at least the first 64 bytes of the failed request */
// enum ofp_queue_op_failed_code {
const (
	OfpQueFailedCodeBadPort = iota /* Invalid port (or port does not exist). */
	OfpQueFailedCodeBadQue         /* Queue does not exist. */
	OfpQueFailedCodeErrPerm        /* Permissions error. */
)

// OfpHelloMsg represents the hello message structure
type OfpHelloMsg struct {
	Header ofpgeneral.OfpHeader
}

// MarshalBinary converts the hello msg fields into byte array
func (hello *OfpHelloMsg) MarshalBinary() (data []byte, err error) {
	return (&hello.Header).MarshalBinary()
}

// UnmarshalBinary transforms the byte array into hello message data
func (hello *OfpHelloMsg) UnmarshalBinary(data []byte) error {
	return (&hello.Header).UnmarshalBinary(data)
}

// OfpPacketInMsg reprensents the packet_in message received by controller
/* Packet received on port (datapath -> controller). */
type OfpPacketInMsg struct {
	Header   ofpgeneral.OfpHeader
	BufferID uint32 /* ID assigned by datapath. */
	TotalLen uint16 /* Full length of frame. */
	InPort   uint16 /* Port on which frame was received. */
	Reason   uint8  /* Reason packet is being sent (one of OFPR_*) */
	//uint8_t pad;
	Data []byte /* Ethernet frame, halfway through 32-bit word,
	   so the IP header is 32-bit aligned.  The
	   amount of data is inferred from the length
	   field in the header.  Because of padding,
	   offsetof(struct ofp_packet_in, data) ==
	   sizeof(struct ofp_packet_in) - 2. */
}

// MarshalBinary converts the packet in msg fields into byte array
func (in *OfpPacketInMsg) MarshalBinary() (data []byte, err error) {
	data = make([]byte, in.TotalLen)
	headerData, err := (&in.Header).MarshalBinary()
	copy(data, headerData)
	binary.BigEndian.PutUint32(data[8:12], in.BufferID)
	binary.BigEndian.PutUint16(data[12:14], in.TotalLen)
	binary.BigEndian.PutUint16(data[14:16], in.Inport)
	data[16] = in.Reason
	copy(data[16:], in.Data)
	return data, err
}

// UnmarshalBinary transforms the byte array into packet in message data
func (in *OfpPacketInMsg) UnmarshalBinary(data []byte) error {
	if err := (&in.Header).UnmarshalBinary(data); err != nil {
		return err
	}
	in.BufferID = binary.BigEndian.Uint32(data[8:12])
	in.TotalLen = binary.BigEndian.Uint16(data[12:14])
	in.InPort = binary.BigEndian.Uint16(data[14:16])
	in.Reason = data[16]
	copy(in.Data, data[16:])
	return nil
}

// OfpPacketOutMsg reprensents the packet_out message sent by controller
/* Send packet (controller -> datapath). */
type OfpPacketOutMsg struct {
	Header     ofpgeneral.OfpHeader
	BufferID   uint32            /* ID assigned by datapath (-1 if none). */
	InPort     uint16            /* Packet's input port (OFPP_NONE if none). */
	ActionsLen uint16            /* Size of action array in bytes. */
	Actions    []OfpActionHeader /* Actions. */
	/* uint8_t data[0]; */ /* Packet data.  The length is inferred
	   from the length field in the header.
	   (Only meaningful if buffer_id == -1.) */
}

// MarshalBinary converts the packet out msg fields into byte array
func (out *OfpPacketOutMsg) MarshalBinary() (data []byte, err error) {
	data = make([]byte, out.TotalLen)
	headerData, err := (&out.Header).MarshalBinary()
	copy(data, headerData)
	binary.BigEndian.PutUint32(data[8:12], out.BufferID)
	binary.BigEndian.PutUint16(data[12:14], out.InPort)
	binary.BigEndian.PutUint16(data[14:16], out.ActionsLen)
	actionByteIdx := 16
	for _, action := range out.Actions {

	}
	return data, err
}

// UnmarshalBinary transforms the byte array into packet out message data
func (out *OfpPacketOutMsg) UnmarshalBinary(data []byte) error {
	if err := (&in.Header).UnmarshalBinary(data); err != nil {
		return err
	}
	in.BufferID = binary.BigEndian.Uint32(data[8:12])
	in.TotalLen = binary.BigEndian.Uint16(data[12:14])
	in.InPort = binary.BigEndian.Uint16(data[14:16])
	in.Reason = data[16]
	copy(in.Data, data[16:])
	return nil
}

// OfpErrMsg represents the msg structure of OFPT_ERROR: Error message (datapath -> controller).
type OfpErrMsg struct {
	Header ofpgeneral.OfpHeader

	Type uint16
	Code uint16
	Data []byte /* Variable-length data.  Interpreted based on the type and code. */
}
