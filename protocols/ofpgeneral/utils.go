package ofpgeneral

import (
	"encoding/binary"
	"fmt"
	"io"
)

// GetMessageVersion is used to retrive the version of the msg from byte slice
func GetMessageVersion(msg []byte) (uint8, error) {
	if len(msg) < 4 {
		return 0, fmt.Errorf("The message length %d is smaller than minimum length", len(msg))
	}
	header := &OfpHeader{}
	if err := header.UnmarshalBinary(msg); err != nil {
		return 0, fmt.Errorf("Version retrival failed due to %s", err.Error())
	}
	return header.Version, nil
}

// GetOfpMsgVersion is used to retrieve the version info of the ofp messge
func GetOfpMsgVersion(msg OfpMessage) (uint8, error) {
	b, err := msg.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return GetMessageVersion(b)
}

// UnMarshalFields is used to read the fields value from reader
func UnMarshalFields(reader io.Reader, fields ...interface{}) error {
	for _, f := range fields {
		if err := binary.Read(reader, binary.BigEndian, f); err != nil {
			return err
		}
	}
	return nil
}

// MarshalFields is used to write the value into fields
func MarshalFields(writer io.Writer, fields ...interface{}) error {
	for _, f := range fields {
		if err := binary.Write(writer, binary.BigEndian, f); err != nil {
			return err
		}
	}
	return nil
}
