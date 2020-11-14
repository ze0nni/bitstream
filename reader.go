package bitstream


type BitReader struct {
	Offset     int
	Buff       []byte
	data       byte
	dataOffset uint8
}



func (b *BitReader) Read_bool() (bool, error) {
	value := b.Buff[b.Offset] > 0
	b.Offset++

	return value, nil
}


func (b *BitReader) Read_byte() (byte, error) {
	value := b.Buff[b.Offset]
	b.Offset++

	return value, nil
}

func (b *BitReader) Read_int8() (int8, error) {
	value := int8(b.Buff[b.Offset])
	b.Offset += 1

	return value, nil
}

func (b *BitReader) Read_int16() (int16, error) {
	value := int16(b.Buff[b.Offset])
	value = (value) | int16(b.Buff[b.Offset+1])<<8
	b.Offset += 2

	return value, nil
}


func (b *BitReader) Read_uint32() (uint32, error) {
	value := uint32(b.Buff[b.Offset])
	value = (value) | uint32(b.Buff[b.Offset+1])<<8
	value = (value) | uint32(b.Buff[b.Offset+2])<<16
	value = (value) | uint32(b.Buff[b.Offset+3])<<24
	b.Offset += 4

	return value, nil
}

func (b *BitReader) Skip(numBytes int) {
	b.Offset += numBytes
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
