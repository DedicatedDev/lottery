package types

const (
	// ModuleName defines the module name
	ModuleName = "lottery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lottery"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo/value/"
)

const (
	LotteryCreatedEventType                    = "new-lottery-created"
	LotteryCreatedEventTypeCreator             = "creator"
	LotteryCreatedEventTypeLotteryId           = "lottery-id"
	LotteryCreatedEventTypeLotteryMinBetAmount = "min-bet-amount"
	LotteryCreatedEventTypeLotteryFee          = "fee"
)

const (
	BidCreatedEventType          = "new-bid-created"
	BidCreatedEventTypeCreator   = "creator"
	BidCreatedEventTypeBidId     = "bid-id"
	BidCreatedEventTypeLotteryId = "lottery-id"
	BidCreatedEventTypeBidCount  = "total-bidder"
)
