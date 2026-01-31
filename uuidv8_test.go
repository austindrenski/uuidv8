package uuidv8

import (
	"encoding/hex"
	"testing"
)

func BenchmarkUUIDv8(b *testing.B) {
	for _, bb := range []struct {
		name string
		t    int64
		h    uint64
	}{
		{
			name: "max",
			t:    0x7FFF_FFFF_FFFF_FFFF,
			h:    0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name: "zero",
			t:    0,
			h:    0,
		},
	} {
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				_ = uuidv8(bb.t, bb.h)
			}
		})
	}
}

func TestUUIDv8(t *testing.T) {
	for _, tt := range []struct {
		name     string
		expected string
		t        int64
		h        uint64
	}{
		{
			name:     "zero",
			expected: "00000000-0000-8000-8000-000000000000",
			t:        0,
			h:        0,
		},
		{
			name:     "max",
			expected: "ffffffff-ffff-8fff-bfff-ffffffffffff",
			t:        0x7FFF_FFFF_FFFF_FFFF,
			h:        0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:     "max timestamp, zero hash",
			expected: "ffffffff-ffff-8fff-8000-000000000000",
			t:        0x7FFF_FFFF_FFFF_FFFF,
			h:        0,
		},
		{
			name:     "zero timestamp, max hash",
			expected: "00000000-0000-8000-bfff-ffffffffffff",
			t:        0,
			h:        0xFFFF_FFFF_FFFF_FFFF,
		},
		{
			name:     "small timestamp",
			expected: "00000000-0001-8000-8000-000000000000",
			t:        4096,
			h:        0,
		},
		{
			name:     "small hash",
			expected: "00000000-0000-8000-8000-000000000001",
			t:        0,
			h:        1,
		},
		{
			// From RFC 9562 Appendix B.1:
			// Timestamp: 164555774200000000 (10 ns-steps) = 0x2489E9AD2EE2E00
			// Expected: 2489E9AD-2EE2-8E00-8EC9-32D5F69181C0
			//
			// The 60-bit timestamp 0x2489E9AD2EE2E00:
			//   custom_a (bits 59-12): 0x2489E9AD2EE2
			//   custom_b (bits 11-0):  0xE00
			//
			// For custom_c, we need to find the hash that produces "8EC9-32D5F69181C0"
			// Byte 8 = 0x8E: variant (0b10) + upper 6 bits of hash (0b001110 = 0x0E)
			// Remaining bytes: C9 32 D5 F6 91 81 C0
			//
			// The 62-bit hash reconstructed: 0x0EC932D5F69181C0
			name:     "RFC 9562 Appendix B.1",
			expected: "2489e9ad-2ee2-8e00-8ec9-32d5f69181c0",
			t:        0x2489E9AD2EE2E00,
			h:        0x0EC932D5F69181C0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			u := uuidv8(tt.t, tt.h)

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

			if tt.expected != string(b[:]) {
				t.Fatalf("expected %q, but received %q", tt.expected, b)
			}
		})
	}
}
