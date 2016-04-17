package ofp13

import (
	"bytes"
	"fmt"

	"github.com/kopwei/goof/protocols/ofpgeneral"
)

// OfpSwitchFeatureMsg represents the switch feature message structure
type OfpSwitchFeatureMsg struct {
	Header     ofpgeneral.OfpHeader
	DatapathID uint64 /* Datapath unique ID.  The lower 48-bits are for
	   a MAC address, while the upper 16-bits are
	   implementer-defined. */

	NoOfBuffers uint32 /* Max packets buffered at once. */

	NoOfTables uint8 /* Number of tables supported by datapath. */

	Padding [3]byte /* Align to 64-bits. */

	/* Features. */
	Capabilities uint32 /* Bitmap of support "ofp_capabilities". */
	Actions      uint32 /* Bitmap of supported "ofp_action_type"s. */

	/* Port info.*/
	Ports []OfpPhysPort /* Port definitions.  The number of ports
	   is inferred from the length field in
	   the header. */
}

// UnmarshalBinary transforms the byte array into header data
func (sf *OfpSwitchFeatureMsg) UnmarshalBinary(data []byte) error {
	if len(data) < 28 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &sf.Header, &sf.DatapathID, &sf.NoOfBuffers,
		&sf.NoOfTables, &sf.Padding, &sf.Capabilities, &sf.Actions, &sf.Ports)
}

// MarshalBinary converts the header fields into byte array
func (sf *OfpSwitchFeatureMsg) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, sf.Header, sf.DatapathID, sf.NoOfBuffers,
		sf.NoOfTables, sf.Padding, sf.Capabilities, sf.Actions, sf.Ports); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// OfpSwitchConfigMsg represents the switch config msg structure
type OfpSwitchConfigMsg struct {
	Header      ofpgeneral.OfpHeader
	Flags       uint16 /* OfpConf* flags. */
	MissSendLen uint16 /* Max bytes of new flow that datapath should
	   send to the controller. */
}

// UnmarshalBinary transforms the byte array into header data
func (sc *OfpSwitchConfigMsg) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("The data size %d is not big enough to be decoded", len(data))
	}
	buf := bytes.NewReader(data)
	return ofpgeneral.UnMarshalFields(buf, &sc.Header, &sc.Flags, &sc.MissSendLen)
}

// MarshalBinary converts the header fields into byte array
func (sc *OfpSwitchConfigMsg) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ofpgeneral.MarshalFields(buf, sc.Header, sc.Flags, sc.MissSendLen); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
