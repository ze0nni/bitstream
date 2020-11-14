package bitstream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/bitstream"
)

func Test_writeZeroBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(false)
	w.Flush()

	assert.Equal(
		t,
		[]byte{0},
		w.Buff,
	)
}

func Test_writeOneBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(true)
	w.Flush()

	assert.Equal(
		t,
		[]byte{1},
		w.Buff,
	)
}

func Test_write2ZeroBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(false)
	w.Write_bool(false)
	w.Flush()

	assert.Equal(
		t,
		[]byte{0},
		w.Buff,
	)
}

func Test_write2OneoBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(true)
	w.Write_bool(true)
	w.Flush()

	assert.Equal(
		t,
		[]byte{3},
		w.Buff,
	)
}

func Test_write16ZeroBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	for i := 0; i < 16; i++ {
		w.Write_bool(false)
	}
	w.Flush()

	assert.Equal(
		t,
		[]byte{0, 0},
		w.Buff,
	)
}

func Test_write8One8ZerBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	for i := 0; i < 8; i++ {
		w.Write_bool(true)
	}
	for i := 0; i < 8; i++ {
		w.Write_bool(false)
	}

	w.Flush()

	assert.Equal(
		t,
		[]byte{255, 0},
		w.Buff,
	)
}

func Test_WriteBytes(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_byte(1)
	w.Write_byte(2)
	w.Write_byte(34)
	w.Flush()

	assert.Equal(
		t,
		[]byte{1, 2, 34},
		w.Buff,
	)
}

func Test_MixOneOneAnd255(t *testing.T) {
	w := &bitstream.BitWriter{}
	w.Write_bool(true)
	w.Write_bool(true)
	w.Write_byte(255)
	w.Flush()

	assert.Equal(
		t,
		[]byte{255, 3},
		w.Buff,
	)
}
