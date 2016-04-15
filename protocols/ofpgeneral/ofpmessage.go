package ofpgeneral

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
