package bitstream

import (
	"bytes"
)

type BitWriter struct {
	Buff []byte
}

type BitReader struct {
	Offset     int
	Buff       []byte
	data       byte
	dataOffset uint8
}

var buff bytes.Buffer

func (b *BitWriter) Write_bool(value bool) error {
	if value {
		b.Buff = append(b.Buff, 1)
	} else {
		b.Buff = append(b.Buff, 0)
	}
	return nil
}

func (b *BitReader) Read_bool() (bool, error) {
	value := b.Buff[b.Offset] > 0
	b.Offset++

	return value, nil
}

func (b *BitWriter) Write_byte(value byte) error {
	b.Buff = append(b.Buff, value)
	return nil
}

func (b *BitReader) Read_byte() (byte, error) {
	value := b.Buff[b.Offset]
	b.Offset++

	return value, nil
}

func (b *BitWriter) Write_int8(value int8) error {
	b.Buff = append(b.Buff, byte(value))
	return nil
}

func (b *BitReader) Read_int8() (int8, error) {
	value := int8(b.Buff[b.Offset])
	b.Offset += 1

	return value, nil
}

func (b *BitWriter) Write_int16(value int16) error {
	b.Buff = append(b.Buff, byte(value>>0))
	b.Buff = append(b.Buff, byte(value>>8))
	return nil
}

func (b *BitReader) Read_int16() (int16, error) {
	value := int16(b.Buff[b.Offset])
	value = (value) | int16(b.Buff[b.Offset+1])<<8
	b.Offset += 2

	return value, nil
}

func (b *BitWriter) Write_uint32(value uint32) error {
	b.Buff = append(b.Buff, byte(value>>0))
	b.Buff = append(b.Buff, byte(value>>8))
	b.Buff = append(b.Buff, byte(value>>16))
	b.Buff = append(b.Buff, byte(value>>24))

	return nil
}

func (b *BitReader) Read_uint32() (uint32, error) {
	value := uint32(b.Buff[b.Offset])
	value = (value) | uint32(b.Buff[b.Offset+1])<<8
	value = (value) | uint32(b.Buff[b.Offset+2])<<16
	value = (value) | uint32(b.Buff[b.Offset+3])<<24
	b.Offset += 4

	return value, nil
}

func (b *BitWriter) Write_float32(value float32) error {

	return nil
}

func (b *BitReader) Read_float32() (float32, error) {

	return 0, nil
}

func (b *BitReader) Skip(numBytes int) {
	b.Offset += numBytes
}

func (b *BitWriter) Flush() error {
	return nil
}

func (b *BitWriter) Reset() {
	b.ResetTo(b.Buff[:0])
}

func (b *BitWriter) ResetTo(buff []byte) {
	b.Buff = buff
}

func (b *BitWriter) ResetToSize(size int) {
	if len(b.Buff) < size {
		b.Buff = append(b.Buff, make([]byte, size-len(b.Buff))...)
	} else {
		b.Buff = b.Buff[:size]
	}
}

func (b *BitReader) Reset(buff []byte) {
	b.Offset = 0
	b.Buff = buff
	b.data = 0
	b.dataOffset = 0
}

func (b *BitReader) EOF() bool {
	if b.Offset+1 >= len(b.Buff) {
		return true
	}

	return false
}
