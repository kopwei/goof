package ofp10

import "github.com/kopwei/goof/protocols/ofpgeneral"

// OfpSwitchFeatureMsg represents the switch feature message structure
type OfpSwitchFeatureMsg struct {
	Header     ofpgeneral.OfpHeader
	DatapathID uint64 /* Datapath unique ID.  The lower 48-bits are for
	   a MAC address, while the upper 16-bits are
	   implementer-defined. */

	NoOfBuffers uint32 /* Max packets buffered at once. */

	NoOfTables uint8 /* Number of tables supported by datapath. */

	//uint8_t pad[3];         /* Align to 64-bits. */

	/* Features. */
	Capabilities uint32 /* Bitmap of support "ofp_capabilities". */
	Actions      uint32 /* Bitmap of supported "ofp_action_type"s. */

	/* Port info.*/
	Ports []OfpPhysPort /* Port definitions.  The number of ports
	   is inferred from the length field in
	   the header. */
}

// OfpSwitchConfigMsg represents the switch config msg structure
type OfpSwitchConfigMsg struct {
	Header      ofpgeneral.OfpHeader
	flags       uint16 /* OfpConf* flags. */
	MissSendLen uint16 /* Max bytes of new flow that datapath should
	   send to the controller. */
}
