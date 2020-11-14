package bitstream

import "errors"

type BitReader struct {
	Offset int
	Buff   []byte

	b  byte
	bi int
}

func (b *BitReader) Read_bool() (bool, error) {
	if 0 == b.bi {
		if b.Offset >= len(b.Buff) {
			return false, errors.New("EOF")
		}

		b.b = b.Buff[b.Offset]
		b.Offset++
	}

	value := 1 == ((b.b >> b.bi) & 1)
	if b.bi == 7 {
		b.bi = 0
	} else {
		b.bi++
	}

	return value, nil
}

func (b *BitReader) Read_byte() (byte, error) {
	if b.Offset >= len(b.Buff) {
		return 0, errors.New("EOF")
	}

	if 0 == b.bi {
		value := b.Buff[b.Offset]
		b.Offset++

		return value, nil
	}
	b0 := b.b >> b.bi
	b.b = b.Buff[b.Offset]
	b1 := b.b << (8 - b.bi)
	b.Offset++

	value := b0 | b1

	return value, nil
}

func (b *BitReader) Read_int8() (int8, error) {
	v, err := b.Read_byte()
	return int8(v), err
}

func (b *BitReader) Read_unt8() (uint8, error) {
	v, err := b.Read_byte()
	return uint8(v), err
}

func (b *BitReader) Read_int16() (int16, error) {
	b0, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b1, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	return int16(b0) | int16(b1)<<8, nil
}

func (b *BitReader) Read_uint16() (uint16, error) {
	b0, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b1, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	return uint16(b0) | uint16(b1)<<8, nil
}

func (b *BitReader) Read_int32() (int32, error) {
	b0, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b1, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b2, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b3, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	return int32(b0) | int32(b1)<<8 | int32(b2)<<16 | int32(b3)<<24, nil
}

func (b *BitReader) Read_uint32() (uint32, error) {
	b0, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b1, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b2, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	b3, err := b.Read_byte()
	if nil != err {
		return 0, err
	}
	return uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24, nil
}

func (b *BitReader) Skip(numBytes int) {
	b.Offset += numBytes
}

func (b *BitReader) Reset(buff []byte) {
	b.Offset = 0
	b.Buff = buff

	b.b = 0
	b.bi = 0
}

func (b *BitReader) EOF() bool {
	if b.Offset+1 >= len(b.Buff) {
		return true
	}

	return false
}
