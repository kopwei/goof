package ofp10

import "github.com/kopwei/goof/protocols/ofpgeneral"

// enum ofp_queue_properties {
const (
	OfpQueNone    = iota /* No property defined for queue (default). */
	OfpQueMinRate        /* Minimum datarate guaranteed. */
	/* Other types should be added here
	 * (i.e. max rate, precedence, etc). */
)

// OfpQueuePropHeader represents the header structure of common description for a queue.
type OfpQueuePropHeader struct {
	Property uint16  /* One of OFPQT_. */
	Len      uint16  /* Length of property, including this header. */
	Paddint  [4]byte /* 64-bit alignemnt. */
}

// OfpQueuePropMinRate represents the min-Rate queue property description.
type OfpQueuePropMinRate struct {
	PropHeader OfpQueuePropHeader /* prop: OFPQT_MIN, len: 16. */
	Rate       uint16             /* In 1/10 of a percent; >1000 -> disabled. */
	Padding    [6]byte            /* 64-bit alignment */
}

// OfpPacketQueue represents the full description for a queue.
type OfpPacketQueue struct {
	QueueID    uint32               /* id for the specific queue. */
	Len        uint16               /* Length in bytes of this queue desc. */
	Padding    [2]byte              /* 64-bit alignment. */
	Properties []OfpQueuePropHeader /* List of properties. */
}

// OfpQueueGetConfReqMsg represents the query msg for port queue configuration.
type OfpQueueGetConfReqMsg struct {
	Header  ofpgeneral.OfpHeader
	Port    uint16  /* Port to be queried. Should refer to a valid physical port (i.e. < OFPP_MAX) */
	Padding [2]byte /* 32-bit alignment. */
}

// OfpQueueGetConfReplyMsg represents queue configuration for a given port.
type OfpQueueGetConfReplyMsg struct {
	Header   ofpgeneral.OfpHeader
	Port     uint16
	Paddingt [6]byte
	Queues   []OfpPacketQueue /* List of configured queues. */
}
