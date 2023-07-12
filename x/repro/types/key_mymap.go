package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MymapKeyPrefix is the prefix to retrieve all Mymap
	MymapKeyPrefix = "Mymap/value/"
)

// MymapKey returns the store key to retrieve a Mymap from the index fields
func MymapKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
