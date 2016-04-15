package ofp10

import "github.com/kopwei/goof/protocols/ofpgeneral"

// OfpMessageParser is the message parser implementation
type OfpMessageParser struct {
}

// ParseMsg is used to convert bytes into ofp message
func (p *OfpMessageParser) ParseMsg(b []byte) (ofpgeneral.OfpMessage, error) {
	return nil, nil
}
