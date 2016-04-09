package ofp15

const (
	// Version is the value of version byte in ofp header
	Version = uint8(0x06)
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

	/* Controller role change request messages */
	OfpTypeRoleRequest = 2 + iota
	OfpTypeRoleReply

	/* Asynchronous message configuration */
	OfpTypeGetAsyncRequest
	OfpTypeGetAsyncReply
	OfpTypeSetAsync

	/* Meters and rate limiters configuration messages. */
	OfpTypeMeterMod

	/* Controller role change event messages */
	OfpTypeRoleStatus

	/* Asynchronous messages */
	OfpTypeTableStatus

	/* Request forwarding by the switch */
	OfpTypeRequestForward

	/* Bundle operations (multiple messages as a single operation) */
	OfpTypeBundleControl
	OfpTypeBundleAddMessage

	/* Controller Status async message */
	OfpTypeControllerStatus
)
