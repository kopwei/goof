package ofp10

import "github.com/kopwei/goof/protocols/ofpgeneral"

//enum ofp_stats_types {
const (
	/* Description of this OpenFlow switch.
	 * The request body is empty.
	 * The reply body is struct ofp_desc_stats. */
	OfpStatsTypeDesc = iota

	/* Individual flow statistics.
	 * The request body is struct ofp_flow_stats_request.
	 * The reply body is an array of struct ofp_flow_stats. */
	OfpStatsTypeFlow

	/* Aggregate flow statistics.
	 * The request body is struct ofp_aggregate_stats_request.
	 * The reply body is struct ofp_aggregate_stats_reply. */
	OfpStatsTypeAggregate

	/* Flow table statistics.
	 * The request body is empty.
	 * The reply body is an array of struct ofp_table_stats. */
	OfpStatsTypeTable

	/* Physical port statistics.
	 * The request body is struct ofp_port_stats_request.
	 * The reply body is an array of struct ofp_port_stats. */
	OfpStatsTypePort

	/* Queue statistics for a port
	 * The request body defines the port
	 * The reply body is an array of struct ofp_queue_stats */
	OfpStatsTypeQueue

	/* Vendor extension.
	 * The request and reply bodies begin with a 32-bit vendor ID, which takes
	 * the same form as in "struct ofp_vendor_header".  The request and reply
	 * bodies are otherwise vendor-defined. */
	OfpStatsTypeVendor = 0xffff
)

// enum ofp_stats_reply_flags {
const (
	OfpStatsReplyMore = iota /* More replies to follow. */
)

const (
	descStrLen   = 256
	serialNumLen = 32
)

// OfpStatsReqMsg represents the structure of Stats request msg
type OfpStatsReqMsg struct {
	Header ofpgeneral.OfpHeader
	Type   uint16 /* One of the OFPST_* constants. */
	Flags  uint16 /* OFPSF_REQ_* flags (none yet defined). */
	Body   []byte /* Body of the request. */
}

// OfpStatsReplyMsg represents the structure of stats reply msg
type OfpStatsReplyMsg struct {
	Header ofpgeneral.OfpHeader
	Type   uint16 /* One of the OFPST_* constants. */
	Flags  uint16 /* OFPSF_REPLY_* flags. */
	Body   []byte /* Body of the reply. */
}

// OfpDescStats represents the structure of descriptive stats
type OfpDescStats struct {
	ManufacurerDesc [descStrLen]byte   /* Manufacturer description. */
	HwDesc          [descStrLen]byte   /* Hardware description. */
	SwDesc          [descStrLen]byte   /* Software description. */
	SerialNum       [serialNumLen]byte /* Serial number. */
	DatapathDesc    [descStrLen]byte   /* Human readable description of datapath. */
}

// OfpFlowStatsReq represents the structure body for ofp_stats_request of type OFPST_FLOW.
type OfpFlowStatsReq struct {
	Match   OfpMatch // Fields to match.
	TableID uint8    //  ID of table to read (from ofp_table_stats), 0xff for all tables or 0xfe for emergency.
	//uint8_t pad;              /* Align to 32 bits. */
	OutPort uint16 // Require matching entries to include this as an output port.  A value of OFPP_NONE indicates no restriction.
}

// OfpFlowStats represents the structure body of reply to OFPST_FLOW request.
type OfpFlowStats struct {
	Length  uint16 /* Length of this entry. */
	TableID uint8  /* ID of table flow came from. */
	//uint8_t pad;
	Match           OfpMatch /* Description of fields. */
	DurationSec     uint32   /* Time flow has been alive in seconds. */
	DurationNanoSec uint32   /* Time flow has been alive in nanoseconds beyond
	   duration_sec. */
	Priority uint16 /* Priority of the entry. Only meaningful
	   when this is not an exact-match entry. */
	IdleTimeout uint16 /* Number of seconds idle before expiration. */
	HardTimeout uint16 /* Number of seconds before expiration. */
	//uint8_t pad2[6];          /* Align to 64-bits. */
	Cookie      uint64            /* Opaque controller-issued identifier. */
	PacketCount uint64            /* Number of packets in flow. */
	ByteCount   uint64            /* Number of bytes in flow. */
	Actions     []OfpActionHeader /* Actions. */
}

// OfpAggStatsRequest represents the structure body for ofp_stats_request of type OFPST_AGGREGATE.
type OfpAggStatsRequest struct {
	Match   OfpMatch /* Fields to match. */
	TableID uint8    /* ID of table to read (from ofp_table_stats) 0xff for all tables or 0xfe for emergency. */
	//uint8_t pad;              /* Align to 32 bits. */
	OutPort uint16 /* Require matching entries to include this as an output port.  A value of OFPP_NONE
	   indicates no restriction. */
}

// OfpAggStatsReply represents the structure body of reply to OFPST_AGGREGATE request. */
type OfpAggStatsReply struct {
	PacketCount uint64 /* Number of packets in flows. */
	ByteCount   uint64 /* Number of bytes in flows. */
	FlowCount   uint32 /* Number of flows. */
	//uint8_t pad[4];           /* Align to 64 bits. */
}

// OfpTableStats represents the structure body of reply to OFPST_TABLE request.
type OfpTableStats struct {
	TableID uint8 // Identifier of table.  Lower numbered tables are consulted first.
	//uint8_t pad[3];          /* Align to 32-bits. */
	Name         [32]byte
	WildCards    uint32 /* Bitmap of OFPFW_* wildcards that are supported by the table. */
	MaxEntries   uint32 /* Max number of entries supported. */
	ActiveCount  uint32 /* Number of active entries. */
	LookupCount  uint64 /* Number of packets looked up in table. */
	MatchedCount uint64 /* Number of packets that hit table. */
}

// OfpPortStatsRequest represents structure body for ofp_stats_request of type OFPST_PORT.
type OfpPortStatsRequest struct {
	// PortNo is the OFPST_PORT message must request statistics
	// either for a single port (specified in
	// port_no) or for all ports (if port_no ==
	// OFPP_NONE).
	PortNo uint16
	//uint8_t pad[6];
}

// OfpPortStats represents the structure body of reply to OFPST_PORT request. If a counter is unsupported, set
// the field to all ones.
type OfpPortStats struct {
	PortNo uint16
	//uint8_t pad[6];          /* Align to 64-bits. */
	RxPackets uint64 /* Number of received packets. */
	TxPackets uint64 /* Number of transmitted packets. */
	RxBytes   uint64 /* Number of received bytes. */
	TxBytes   uint64 /* Number of transmitted bytes. */
	RxDropped uint64 /* Number of packets dropped by RX. */
	TxDropped uint64 /* Number of packets dropped by TX. */
	RxErrors  uint64 /* Number of receive errors.  This is a super-set
	   of more specific receive errors and should be
	   greater than or equal to the sum of all
	   rx_*_err values. */
	TxErrors uint64 /* Number of transmit errors.  This is a super-set
	   of more specific transmit errors and should be
	   greater than or equal to the sum of all
	   tx_*_err values (none currently defined.) */
	RxFrameErr uint64 /* Number of frame alignment errors. */
	RxOverErr  uint64 /* Number of packets with RX overrun. */
	RxCrcErr   uint64 /* Number of CRC errors. */
	Collisions uint64 /* Number of collisions. */
}

// OfpVendorHeader represents the header structure of Vendor extension.
type OfpVendorHeader struct {
	Header ofpgeneral.OfpHeader /* Type OFPT_VENDOR. */
	Vendor uint32               /* Vendor ID:
	 * - MSB 0: low-order bytes are IEEE OUI.
	 * - MSB != 0: defined by OpenFlow
	 *   consortium. */
	/* Vendor-defined arbitrary additional data. */
}

// OfpQueueStatsReq represents the ofp queue stats query structure
type OfpQueueStatsReq struct {
	PortNo uint16 /* All ports if OFPT_ALL. */
	//uint8_t pad[2];          /* Align to 32-bits. */
	QueueID uint32 /* All queues if OFPQ_ALL. */
}

// OfpQueueStats represents the queue stats info
type OfpQueueStats struct {
	PortNo uint16
	//uint8_t pad[2];          /* Align to 32-bits. */
	QueueID   uint32 /* Queue i.d */
	TxBytes   uint64 /* Number of transmitted bytes. */
	TxPackets uint64 /* Number of transmitted packets. */
	TxErrors  uint64 /* Number of packets dropped due to overrun. */
}
