package bitstream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/bitstream"
)

func Test_Mixed(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(true) // b0
	w.Write_byte(1)
	w.Write_byte(22)
	w.Write_byte(44)

	w.Write_bool(false) // b1
	w.Write_int8(-55)
	w.Write_int8(56)

	w.Write_bool(true) // b2
	w.Write_uint8(200)

	w.Write_bool(false) //b3
	w.Write_int16(1111)
	w.Write_int16(2222)
	w.Write_int16(-3333)

	w.Write_bool(true) // b4
	w.Write_uint16(65432)

	w.Write_bool(false) //b5
	w.Write_int32(0xabcdef)
	w.Write_int32(int32(-0xabcdef))

	w.Write_uint32(0xffabcdef)

	r := &bitstream.BitReader{}
	r.Reset(w.Buff)

	b0, _ := r.Read_bool()
	assert.Equal(t, true, b0, "b0")

	bt1, _ := r.Read_byte()
	assert.Equal(t, byte(1), bt1, "bt1")

	bt2, _ := r.Read_byte()
	assert.Equal(t, byte(22), bt2, "bt2")

	bt3, _ := r.Read_byte()
	assert.Equal(t, byte(44), bt3, "bt3")

	b1, _ := r.Read_bool()
	assert.Equal(t, false, b1, "b1")

	int8_m55, _ := r.Read_int8()
	assert.Equal(t, int8(-55), int8_m55, "int8_m55")

	int8_56, _ := r.Read_int8()
	assert.Equal(t, int8(56), int8_56, "int8_56")

	b2, _ := r.Read_bool()
	assert.Equal(t, true, b2, "b2")

	uint8_200, _ := r.Read_uint8()
	assert.Equal(t, uint8(200), uint8_200, "uint8_200")

	b3, _ := r.Read_bool()
	assert.Equal(t, false, b3, "b3")

	int16_1111, _ := r.Read_int16()
	assert.Equal(t, int16(1111), int16_1111, "int16_1111")

	int16_2222, _ := r.Read_int16()
	assert.Equal(t, int16(2222), int16_2222, "int16_2222")

	int16_m3333, _ := r.Read_int16()
	assert.Equal(t, int16(-3333), int16_m3333, "int16_m3333")

	b4, _ := r.Read_bool()
	assert.Equal(t, true, b4, "b4")

	uint16_65432, _ := r.Read_uint16()
	assert.Equal(t, uint16(65432), uint16_65432, "uint16_65432")

	b5, _ := r.Read_bool()
	assert.Equal(t, false, b5, "b5")

	int32_0xabcdef, _ := r.Read_int32()
	assert.Equal(t, int32(0xabcdef), int32_0xabcdef)

	int32_m0xabcdef, _ := r.Read_int32()
	assert.Equal(t, int32(-0xabcdef), int32_m0xabcdef)

	uint32_0xffabcdef, _ := r.Read_uint32()
	assert.Equal(t, uint32_0xffabcdef, uint32_0xffabcdef)

	w.Flush()
}
