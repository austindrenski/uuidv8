package uuidv8_test

import (
	"encoding/hex"
	"testing"

	"go.austindrenski.io/uuidv8"
)

func BenchmarkUUIDv8(b *testing.B) {
	for _, bb := range []struct {
		name      string
		timestamp int64
		hash      uint64
	}{
		{
			name:      "zero timestamp, zero hash",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "max timestamp, max hash",
			timestamp: 0x7FFF_FFFF_FFFF_FFFF,
			hash:      0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:      "max timestamp, zero hash",
			timestamp: 0x7FFF_FFFF_FFFF_FFFF,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "zero timestamp, max hash",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:      "zero timestamp, min hash",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0x0000_0000_0000_0070,
		},
		{
			name:      "min timestamp, zero hash",
			timestamp: 0x0000_0000_0000_0001,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "RFC 9562 Appendix B.1 (adapted for 1 ns-steps)",
			timestamp: 0x0248_9E9A_D2EE_2E00,
			hash:      0x0EC9_32D5_F691_81C0,
		},
	} {
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				_ = uuidv8.UUIDv8(bb.timestamp, bb.hash)
			}
		})
	}
}

func TestUUIDv8(t *testing.T) {
	for _, tt := range []struct {
		name      string
		expected  string
		timestamp int64
		hash      uint64
	}{
		{
			name:      "zero timestamp, zero hash",
			expected:  "00000000-0000-8000-8000-000000000000",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "max timestamp, max hash",
			expected:  "7fffffff-ffff-8fff-bfff-ffffffffffff",
			timestamp: 0x7FFF_FFFF_FFFF_FFFF,
			hash:      0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:      "max timestamp, zero hash",
			expected:  "7fffffff-ffff-8fff-bc00-000000000000",
			timestamp: 0x7FFF_FFFF_FFFF_FFFF,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "zero timestamp, max hash",
			expected:  "00000000-0000-8000-83ff-ffffffffffff",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:      "zero timestamp, min hash",
			expected:  "00000000-0000-8000-8000-000000000001",
			timestamp: 0x0000_0000_0000_0000,
			hash:      0x0000_0000_0000_0070,
		},
		{
			name:      "min timestamp, zero hash",
			expected:  "00000000-0000-8000-8400-000000000000",
			timestamp: 0x0000_0000_0000_0001,
			hash:      0x0000_0000_0000_0000,
		},
		{
			name:      "RFC 9562 Appendix B.1 (adapted for 1 ns-steps)",
			expected:  "02489e9a-d2ee-82e0-803b-24cb57da4607",
			timestamp: 0x0248_9E9A_D2EE_2E00,
			hash:      0x0EC9_32D5_F691_81C0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			u := uuidv8.UUIDv8(tt.timestamp, tt.hash)

			if actual := format(u); tt.expected != actual {
				t.Errorf("expected %q, but received %q", tt.expected, actual)
			}

			if actual := timestamp(u); tt.timestamp != actual {
				t.Errorf("expected %q, but received %q", tt.timestamp, actual)
			}

			if actual := u[8] & 0b1100_0000 >> 6; byte(0b10) != actual {
				t.Errorf("expected %v, but received %v", byte(0b10), actual)
			}

			if actual := u[6] & 0b1111_0000 >> 4; byte(0b1000) != actual {
				t.Errorf("expected %v, but received %v", byte(0b1000), actual)
			}
		})
	}
}

func format(u [16]byte) string {
	var b [36]byte

	hex.Encode(b[:], u[0:4])
	b[8] = '-'
	hex.Encode(b[9:13], u[4:6])
	b[13] = '-'
	hex.Encode(b[14:18], u[6:8])
	b[18] = '-'
	hex.Encode(b[19:23], u[8:10])
	b[23] = '-'
	hex.Encode(b[24:36], u[10:16])

	return string(b[:])
}

func timestamp(u [16]byte) int64 {
	var t int64

	t |= int64(u[0]) << 56
	t |= int64(u[1]) << 48
	t |= int64(u[2]) << 40
	t |= int64(u[3]) << 32
	t |= int64(u[4]) << 24
	t |= int64(u[5]) << 16
	t |= int64(u[6]&0b0000_1111) << 12
	t |= int64(u[7]) << 4
	t |= int64(u[8]&0b0011_1100) >> 2

	return t
}
