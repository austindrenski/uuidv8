// Package uuidv8 is a reference implementation to construct
// UUIDs using 64-bit timestamps based on [RFC 9562 §5.8].
//
// Whereas [RFC 9562 §B.1] provides an illustrative example to store
// timestamps in 10-ns steps with 62 bits of additional data, this
// implementation encodes the full timestamp in 1-ns steps leaving
// room for 58 bits of additional data.
//
// Whether this trade-off is suitable will depend on each individual
// use case. The collision risk depends on many factors, including
// the frequency of UUID generation and the entropy of the remaining
// 58 bits.
//
// One example where this this trade-off can be advantageous is to
// construct UUIDs for telemetry data where timestamps are natively
// encoded as nanoseconds since the Unix epoch.
//
// In such cases, UUIDv8 can be used to construct deterministic
// identifiers for telemetry data based on the signal timestamp
// and a content-based hash (e.g. xxh3) of the signal payload.
//
// In such cases, encoding the full timestamp greatly narrows the
// window for collisions in the remaining 58 bits, while retaining
// the lexicographical ordering properties of the native timestamp.
//
// [5.8.  UUID Version 8][RFC 9562 §5.8]
//
//	UUIDv8 provides a format for experimental or vendor-specific use
//	cases.  The only requirement is that the variant and version bits
//	MUST be set as defined in Sections 4.1 and 4.2.  UUIDv8's uniqueness
//	will be implementation specific and MUST NOT be assumed.
//
//	The only explicitly defined bits are those of the version and variant
//	fields, leaving 122 bits for implementation-specific UUIDs.  To be
//	clear, UUIDv8 is not a replacement for UUIDv4 (Section 5.4) where all
//	122 extra bits are filled with random data.
//
//	Some example situations in which UUIDv8 usage could occur:
//
//	*  An implementation would like to embed extra information within the
//	   UUID other than what is defined in this document.
//
//	*  An implementation has other application and/or language
//	   restrictions that inhibit the use of one of the current UUIDs.
//
//	Appendix B provides two illustrative examples of custom UUIDv8
//	algorithms to address two example scenarios.
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           custom_a                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|          custom_a             |  ver  |       custom_b        |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|var|                       custom_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           custom_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
//	                Figure 12: UUIDv8 Field and Bit Layout
//
//	custom_a:
//	   The first 48 bits of the layout that can be filled as an
//	   implementation sees fit.  Occupies bits 0 through 47 (octets 0-5).
//
//	ver:
//	   The 4-bit version field as defined by Section 4.2, set to 0b1000
//	   (8).  Occupies bits 48 through 51 of octet 6.
//
//	custom_b:
//	   12 more bits of the layout that can be filled as an implementation
//	   sees fit.  Occupies bits 52 through 63 (octets 6-7).
//
//	var:
//	   The 2-bit variant field as defined by Section 4.1, set to 0b10.
//	   Occupies bits 64 and 65 of octet 8.
//
//	custom_c:
//	   The final 62 bits of the layout immediately following the var
//	   field to be filled as an implementation sees fit.  Occupies bits
//	   66 through 127 (octets 8-15).
//
// [RFC 9562 §5.8]: https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-8
// [RFC 9562 §B.1]: https://www.rfc-editor.org/rfc/rfc9562#name-example-of-a-uuidv8-value-t
package uuidv8 // import "go.austindrenski.io/uuidv8"
