package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredLotteryKeyPrefix is the prefix to retrieve all StoredLottery
	StoredLotteryKeyPrefix = "StoredLottery/value/"
)

// StoredLotteryKey returns the store key to retrieve a StoredLottery from the index fields
func StoredLotteryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
