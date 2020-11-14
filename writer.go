package bitstream

type BitWriter struct {
	b  byte
	bi int

	Buff []byte
}

func (b *BitWriter) Write_bool(value bool) error {
	if value {
		b.b = b.b | (1 << b.bi)
	}
	if b.bi < 7 {
		b.bi++
	} else {
		b.Buff = append(b.Buff, b.b)
		b.b = 0
		b.bi = 0
	}

	return nil
}

func (b *BitWriter) Write_byte(value byte) error {
	if 0 == b.bi {
		b.Buff = append(b.Buff, value)
	} else {
		b.b = b.b | (value << b.bi)
		b.Buff = append(b.Buff, b.b)
		b.b = (value >> (8 - b.bi))
	}
	return nil
}

func (b *BitWriter) Write_int8(value int8) error {
	return b.Write_byte(byte(value))
}

func (b *BitWriter) Write_uint8(value uint8) error {
	return b.Write_byte(byte(value))
}

func (b *BitWriter) Write_int16(value int16) error {
	return b.Write_byte(byte(value >> 0))
	return b.Write_byte(byte(value >> 8))
	return nil
}

func (b *BitWriter) Write_uint16(value uint16) error {
	return b.Write_byte(byte(value >> 0))
	return b.Write_byte(byte(value >> 8))
	return nil
}

func (b *BitWriter) Write_int32(value int32) error {
	return b.Write_byte(byte(value >> 0))
	return b.Write_byte(byte(value >> 8))
	return b.Write_byte(byte(value >> 16))
	return b.Write_byte(byte(value >> 24))

	return nil
}

func (b *BitWriter) Write_uint32(value uint32) error {
	return b.Write_byte(byte(value >> 0))
	return b.Write_byte(byte(value >> 8))
	return b.Write_byte(byte(value >> 16))
	return b.Write_byte(byte(value >> 24))

	return nil
}

func (b *BitWriter) Flush() error {
	if 0 != b.bi {
		b.Buff = append(b.Buff, b.b)
	}
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
