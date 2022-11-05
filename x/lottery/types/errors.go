package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrFinishedLottery  = sdkerrors.Register(ModuleName, 1100, "lottery is finished")
	ErrInvalidBidder    = sdkerrors.Register(ModuleName, 1101, "invalid bidder")
	ErrInvalidLottery   = sdkerrors.Register(ModuleName, 1102, "invalid lottery")
	ErrInvalidLotteryId = sdkerrors.Register(ModuleName, 1103, "invalid lottery id")
	ErrInvalidBetAmount = sdkerrors.Register(ModuleName, 1104, "invalid bet amount")
	ErrInvalidFeeAmount = sdkerrors.Register(ModuleName, 1105, "invalid fee amount")
	ErrInvalidBid       = sdkerrors.Register(ModuleName, 1106, "invalid bid")
)
