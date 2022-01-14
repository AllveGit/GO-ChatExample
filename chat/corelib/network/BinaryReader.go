package network

import (
	"chat/corelib/errs"
	"encoding/binary"
)

type BinaryReader struct {
	Bytes []byte
	pos   int
}

func NewBinaryReader(InBytes []byte) *BinaryReader {
	reader := &BinaryReader{
		Bytes: InBytes,
	}

	return reader
}

func (self *BinaryReader) Read(InBytes []byte) (int, error) {
	remain := len(self.Bytes) - self.pos
	want := len(InBytes)

	if remain < want {
		err := errs.New("Read failed. not enough Bytes for read. len: %d, Buffer: %v", len(self.Bytes), self.Bytes)
		return 0, err
	}

	var n int
	if remain < want {
		n = remain
	} else {
		n = want
	}

	copy(InBytes, self.Bytes[self.pos:self.pos+n])
	self.pos = self.pos + n
	return n, nil
}

func (self *BinaryReader) ReadInt8() (int8, error) {
	var v int8
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadInt16() (int16, error) {
	var v int16
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadInt32() (int32, error) {
	var v int32
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadInt64() (int64, error) {
	var v int64
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadUInt8() (uint8, error) {
	var v uint8
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadUInt16() (uint16, error) {
	var v uint16
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadUInt32() (uint32, error) {
	var v uint32
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}

func (self *BinaryReader) ReadUInt64() (uint64, error) {
	var v uint64
	err := binary.Read(self, binary.LittleEndian, &v)
	return v, err
}
