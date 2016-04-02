package ofp10

import (
	"net"

	"github.com/kopwei/goof/protocols/ofpgeneral"
)

// ofp_flow_mod_command
const (
	OfpFlowModCmdAdd          = iota /* New flow. */
	OfpFlowModCmdModify              /* Modify all matching flows. */
	OfpFlowModCmdModifyStrict        /* Modify entry strictly matching wildcards */
	OfpFlowModCmdDelete              /* Delete all matching flows. */
	OfpFlowModCmdDeleteStrict        /* Strictly match wildcards and priority. */
)

// Flow wildcards.
// enum ofp_flow_wildcards {
const (
	OfpFlowWildCardsInPort  = 1 << iota /* Switch input port. */
	OfpFlowWildCardsDLVlan              /* VLAN id. */
	OfpFlowWildCardsDLSrc               /* Ethernet source address. */
	OfpFlowWildCardsDLDst               /* Ethernet destination address. */
	OfpFlowWildCardsDLType              /* Ethernet frame type. */
	OfpFlowWildCardsNWProto             /* IP protocol. */
	OfpFlowWildCardsTPSrc               /* TCP/UDP source port. */
	OfpFlowWildCardsTPDst               /* TCP/UDP destination port. */

	/* IP source address wildcard bit count.  0 is exact match, 1 ignores the
	 * LSB, 2 ignores the 2 least-significant bits, ..., 32 and higher wildcard
	 * the entire field.  This is the *opposite* of the usual convention where
	 * e.g. /24 indicates that 8 bits (not 24 bits) are wildcarded. */
	OfpFlowWildCardsNWSrcShift = 8
	OfpFlowWildCardsNWSrcBits  = 6
	OfpFlowWildCardsNWSrcMask  = ((1 << OfpFlowWildCardsNWSrcBits) - 1) << OfpFlowWildCardsNWSrcShift
	OfpFlowWildCardsNWSrcAll   = 32 << OfpFlowWildCardsNWSrcShift

	/* IP destination address wildcard bit count.  Same format as source. */
	OfpFlowWildCardsNWDstShift = 14
	OfpFlowWildCardsNWDstBits  = 6
	OfpFlowWildCardsNWDstMask  = ((1 << OfpFlowWildCardsNWDstBits) - 1) << OfpFlowWildCardsNWDstShift
	OfpFlowWildCardsNWDstAll   = 32 << OfpFlowWildCardsNWDstShift

	OfpFlowWildCardsDLVlanPCP = 1 << 20 /* VLAN priority. */
	OfpFlowWildCardsNWToS     = 1 << 21 /* IP ToS (DSCP field, 6 bits). */

	/* Wildcard all fields. */
	OfpFlowWildCardsALL = ((1 << 22) - 1)
)

// ofp_flow_mod_flags {
const (
	OfpFlowFlagSendFlowRemove = 1 << iota /* Send flow removed message when flow
	 * expires or is deleted. */
	OfpFlowFlagCheckOverlap /* Check for overlapping entries first. */
	OfpFlowFlagEmergency    /* Remark this is for emergency. */
)

//  ofp_flow_removed_reason
const (
	OfpFlowRemoveReasonIdleTimeout = iota /* Flow idle time exceeded idle_timeout. */
	OfpFlowRemoveReasonHardTimeout        /* Time exceeded hard_timeout. */
	OfpFlowRemoveReasonDelete             /* Evicted by a DELETE flow mod. */
)

// ofp_error_msg 'code' values for OFPET_FLOW_MOD_FAILED.  'data' contains
// at least the first 64 bytes of the failed request. */
// enum ofp_flow_mod_failed_code {
const (
	OfpFlowModFailedAllTablesFull   = iota /* Flow not added because of full tables. */
	OfpFlowModFailedOverlap                /* Attempted to add overlapping flow with CHECK_OVERLAP flag set. */
	OfpFlowModFailedErrPerm                /* Permissions error. */
	OfpFlowModFailedBadEmergTimeout        /* Flow not added because of non-zero idle/hard timeout. */
	OfpFlowModFailedBadCmd                 /* Unknown command. */
	OfpFlowModFailedUnsupported            /* Unsupported action list - cannot process in the order specified. */
)

// OfpMatch are fields to match against flows
type OfpMatch struct {
	Wildcards uint32           /* Wildcard fields. */
	InPort    uint16           /* Input switch port. */
	DLSrc     net.HardwareAddr /* Ethernet source address. */
	DLDst     net.HardwareAddr /* Ethernet destination address. */
	DLVlan    uint16           /* Input VLAN id. */
	DLVlanPCP uint8            /* Input VLAN priority. */
	//uint8_t pad1[1];           /* Align to 64-bits */
	DLType  uint16 /* Ethernet frame type. */
	NWToS   uint8  /* IP ToS (actually DSCP field, 6 bits). */
	NWProto uint8  /* IP protocol or lower 8 bits of
	 * ARP opcode. */
	//uint8_t pad2[2];           /* Align to 64-bits */
	NWSrc net.IP /* IP source address. */
	NWDst net.IP /* IP destination address. */
	TPSrc uint16 /* TCP/UDP source port. */
	TPDst uint16 /* TCP/UDP destination port. */
}

// OfpModFlowMsg represents the structure of flow setup and teardown (controller -> datapath).
type OfpModFlowMsg struct {
	Header ofpgeneral.OfpHeader
	Match  OfpMatch /* Fields to match */
	Cookie uint64   /* Opaque controller-issued identifier. */

	/* Flow actions. */
	Command     uint16 /* One of OFPFC_*. */
	IdleTimeout uint16 /* Idle time before discarding (seconds). */
	HardTimeout uint16 /* Max time before discarding (seconds). */
	Priority    uint16 /* Priority level of flow entry. */
	BufferID    uint32 /* Buffered packet to apply to (or -1).
	   Not meaningful for OFPFC_DELETE*. */
	OutPort uint16 /* For OFPFC_DELETE* commands, require
	   matching entries to include this as an
	   output port.  A value of OFPP_NONE
	   indicates no restriction. */
	Flags   uint16            /* One of OFPFF_*. */
	Actions []OfpActionHeader /* The action length is inferred
	   from the length field in the
	   header. */
}

// OfpFlowRemovedMsg represents the msg structure of flow removed (datapath -> controller).
type OfpFlowRemovedMsg struct {
	Header ofpgeneral.OfpHeader
	Match  OfpMatch /* Description of fields. */
	Cookie uint64   /* Opaque controller-issued identifier. */

	Priority uint16 /* Priority level of flow entry. */
	Reason   uint8  /* One of OFPRR_*. */
	//uint8_t pad[1];           /* Align to 32-bits. */

	DurationSec     uint32 /* Time flow was alive in seconds. */
	DurationNanoSec uint32 /* Time flow was alive in nanoseconds beyond
	   duration_sec. */
	IdleTimeout uint16 /* Idle timeout from original flow mod. */
	//uint8_t pad2[2];          /* Align to 64-bits. */
	PacketCount uint64
	ByteCount   uint64
}
