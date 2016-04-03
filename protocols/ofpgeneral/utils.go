package ofpgeneral

import (
	"encoding/binary"
	"io"
)

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
