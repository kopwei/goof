package ofpgeneral

import "bytes"

// OfpMessage is the general representation of the messages
// transferred between controller and switch
type OfpMessage interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}

// OfpHelloMsg represents the hello message structure
type OfpHelloMsg struct {
	Header OfpHeader
}

// MarshalBinary converts the hello msg fields into byte array
func (hello *OfpHelloMsg) MarshalBinary() ([]byte, error) {
	return (&hello.Header).MarshalBinary()
}

// UnmarshalBinary transforms the byte array into hello message data
func (hello *OfpHelloMsg) UnmarshalBinary(data []byte) error {
	return (&hello.Header).UnmarshalBinary(data)
}

// NewHelloMsg creates a hello message
func NewHelloMsg(version uint8) *OfpHelloMsg {
	header := NewOfpHeader(version)
	return &OfpHelloMsg{Header: *header}
}

// OfpErrMsg represents the msg structure of OFPT_ERROR: Error message (datapath -> controller).
type OfpErrMsg struct {
	Header OfpHeader

	Type uint16
	Code uint16
	Data []byte /* Variable-length data.  Interpreted based on the type and code. */
}

// MarshalBinary converts the packet in msg fields into byte array
func (em *OfpErrMsg) MarshalBinary() ([]byte, error) {
	data := make([]byte, em.Header.Length)
	headerData, err := (&em.Header).MarshalBinary()
	copy(data, headerData)
	buf := new(bytes.Buffer)
	err = MarshalFields(buf, em.Type, em.Code)
	if err != nil {
		return nil, err
	}
	copy(data[4:8], buf.Bytes())
	copy(data[8:], em.Data)
	return data, err
}

// UnmarshalBinary transforms the byte array into packet in message data
func (em *OfpErrMsg) UnmarshalBinary(data []byte) error {
	if err := (&em.Header).UnmarshalBinary(data); err != nil {
		return err
	}
	buf := bytes.NewReader(data[4:8])
	if err := UnMarshalFields(buf, &em.Type, &em.Code); err != nil {
		return err
	}
	copy(em.Data, data[8:])
	return nil
}
