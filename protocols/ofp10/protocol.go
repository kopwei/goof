package ofp10

import "github.com/kopwei/goof/protocols/ofpgeneral"

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

// OfpHelloMsg represents the hello message structure
type OfpHelloMsg struct {
	Header ofpgeneral.OfpHeader
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

// OfpPacketOutMsg reprensents the packet_out message sent by controller
/* Send packet (controller -> datapath). */
type OfpPacketOutMsg struct{
    Header   ofpgeneral.OfpHeader
    BufferID uint32            /* ID assigned by datapath (-1 if none). */
    InPort   uint16           /* Packet's input port (OFPP_NONE if none). */
    ActionsLen uint16         /* Size of action array in bytes. */
    Actions []OfpActionHeader /* Actions. */
    /* uint8_t data[0]; */        /* Packet data.  The length is inferred
                                     from the length field in the header.
                                     (Only meaningful if buffer_id == -1.) */
}