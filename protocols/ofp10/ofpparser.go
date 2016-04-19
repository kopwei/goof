package ofp10

import (
	"errors"

	"github.com/kopwei/goof/protocols/ofpgeneral"
)

// OfpMessageParser is the message parser implementation
type OfpMessageParser struct {
}

// ParseMsg is used to convert bytes into ofp message
func (p *OfpMessageParser) ParseMsg(b []byte) (ofpgeneral.OfpMessage, error) {
	var message ofpgeneral.OfpMessage
	var err error
	switch b[1] {
	case OfpTypeHello:
		message = &ofpgeneral.OfpHelloMsg{}
		message.UnmarshalBinary(b)
	case OfpTypeError:
		message = &ofpgeneral.OfpErrMsg{}
		message.UnmarshalBinary(b)
	case OfpTypeEchoRequest:
		message = &ofpgeneral.OfpHeader{}
		message.UnmarshalBinary(b)
	case OfpTypeEchoReply:
		message = &ofpgeneral.OfpHeader{}
		message.UnmarshalBinary(b)
	default:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	}
	return message, err
}
