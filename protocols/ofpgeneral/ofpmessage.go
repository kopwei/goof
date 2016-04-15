package ofpgeneral

// OfpMessage is the general representation of the messages
// transferred between controller and switch
type OfpMessage interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) error
}
