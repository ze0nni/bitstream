package bitstream_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/bitstream"
)

func Test_readZeroBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(false)
	w.Flush()

	r := &bitstream.BitReader{}
	r.Reset(w.Buff)

	b, err := r.Read_bool()

	assert.NoError(t, err)

	assert.Equal(
		t,
		false,
		b,
	)
}

func Test_readOneBit(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(true)
	w.Flush()

	r := &bitstream.BitReader{}
	r.Reset(w.Buff)

	b, err := r.Read_bool()

	assert.NoError(t, err)

	assert.Equal(
		t,
		true,
		b,
	)
}

func Test_readOneZeroOneBits(t *testing.T) {
	w := &bitstream.BitWriter{}

	w.Write_bool(true)
	w.Write_bool(false)
	w.Write_bool(true)
	w.Flush()

	r := &bitstream.BitReader{}
	r.Reset(w.Buff)

	b0, err := r.Read_bool()
	assert.NoError(t, err, "1")
	assert.Equal(t, true, b0, "1'")

	b1, err := r.Read_bool()
	assert.NoError(t, err, "2")
	assert.Equal(t, false, b1, "2'")

	b2, err := r.Read_bool()
	assert.NoError(t, err, "3")
	assert.Equal(t, true, b2, "3")
}

func Test_read8One8ZeroBits(t *testing.T) {
	w := &bitstream.BitWriter{}

	for i := 0; i < 9; i++ {
		w.Write_bool(true)
	}
	for i := 0; i < 9; i++ {
		w.Write_bool(false)
	}
	w.Flush()

	r := &bitstream.BitReader{}
	r.Reset(w.Buff)

	for i := 0; i < 9; i++ {
		b, err := r.Read_bool()
		assert.NoError(t, err, "1 "+strconv.Itoa(i))
		assert.Equal(t, true, b, "1 "+strconv.Itoa(i))
	}
	for i := 0; i < 9; i++ {
		b, err := r.Read_bool()
		assert.NoError(t, err, "0 "+strconv.Itoa(i))
		assert.Equal(t, false, b, "0 "+strconv.Itoa(i))
	}
}
