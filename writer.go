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
	b.Buff = append(b.Buff, value)
	return nil
}

func (b *BitWriter) Write_int8(value int8) error {
	b.Buff = append(b.Buff, byte(value))
	return nil
}

func (b *BitWriter) Write_int16(value int16) error {
	b.Buff = append(b.Buff, byte(value>>0))
	b.Buff = append(b.Buff, byte(value>>8))
	return nil
}

func (b *BitWriter) Write_uint32(value uint32) error {
	b.Buff = append(b.Buff, byte(value>>0))
	b.Buff = append(b.Buff, byte(value>>8))
	b.Buff = append(b.Buff, byte(value>>16))
	b.Buff = append(b.Buff, byte(value>>24))

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
