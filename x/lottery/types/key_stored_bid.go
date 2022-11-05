package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredBidKeyPrefix is the prefix to retrieve all StoredBid
	StoredBidKeyPrefix = "StoredBid/value/"
)

// StoredBidKey returns the store key to retrieve a StoredBid from the index fields
func StoredBidKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
