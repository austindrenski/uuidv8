# uuidv8

[![Go Reference](https://pkg.go.dev/badge/go.austindrenski.io/uuidv8.svg)](https://pkg.go.dev/go.austindrenski.io/uuidv8)
[![GitHub Actions](https://github.com/austindrenski/uuidv8/actions/workflows/ci.yaml/badge.svg)](https://github.com/austindrenski/uuidv8/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/go.austindrenski.io/uuidv8)](https://goreportcard.com/report/go.austindrenski.io/uuidv8)
[![GitHub Release](https://github.com/austindrenski/uuidv8/releases)](https://img.shields.io/github/v/release/austindrenski/uuidv8?style=for-the-badge)

This is a reference implementation to construct UUIDs using 64-bit timestamps based
on [RFC 9562 §5.8](https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-8).

Whereas [RFC 9562 §B.1](https://www.rfc-editor.org/rfc/rfc9562#name-example-of-a-uuidv8-value-t) provides an illustrative example to store timestamps
in 10-ns steps with 62 bits of additional data, this implementation
encodes the full timestamp in 1-ns steps leaving room for 58 bits of additional data.

Whether this trade-off is suitable will depend on each individual use case. The collision risk depends on many factors, including the frequency of
UUID generation and the entropy of the remaining 58 bits.

One example where this this trade-off can be advantageous is to construct UUIDs for telemetry data where timestamps are natively encoded as
nanoseconds since the Unix epoch.

In such cases, UUIDv8 can be used to construct deterministic identifiers for telemetry data based on the signal timestamp and a content-based hash (
e.g. xxh3) of the signal payload.

In such cases, encoding the full timestamp greatly narrows the window for collisions in the remaining 58 bits, while retaining the lexicographical
ordering properties of the native timestamp.

There are many potential use cases where constructing a UUIDv8 from a timestamp with a specific precision will be useful, and this implementation can
be adapted to support other timestamp precisions (e.g. 10-ns, 1-ms, 1-s) by modifying the way the timestamp is encoded into the UUIDv8 layout.

___[This is free and unencumbered software released into the public domain](./LICENSE), and as such, you are welcome to use this packackge directly,
to fork this repo, and to copy/paste this code without acknowledgement as you see fit.___

_(That said, some acknowledgement is always appreciated, so consider starring this repository if you find it useful!)_

## Installation

```sh
go get go.austindrenski.io/uuidv8
```

## License

This is free and unencumbered software released into the public domain.

See [LICENSE](./LICENSE) for details.
