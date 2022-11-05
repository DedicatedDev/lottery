package types

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (bid StoredBid) GetBidderAddress() (sdk.AccAddress, error) {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		return nil, err
	}
	return bidder, nil
}

// func (bid StoredBid) GetBidIndex() (string, error) {
// 	index := sha512.New()
// 	sum := strconv.Itoa(int(bid.LotteryId)) + bid.Bidder
// 	index.Write([]byte(sum))
// 	indexHashString := base64.StdEncoding.EncodeToString(index.Sum(nil))
// 	return indexHashString, nil
// }

func GenValidBidIndex(bidder string, id uint64) (string, error) {
	_, err := sdk.AccAddressFromBech32(bidder)
	if err != nil {
		return "", err
	}
	index := sha512.New()
	sum := strconv.Itoa(int(id)) + bidder
	index.Write([]byte(sum))
	indexHashString := base64.StdEncoding.EncodeToString(index.Sum(nil))
	return indexHashString, nil
}

func (lottery *StoredLottery) ValidateLottery() error {
	if lottery.Index == "" {
		return ErrInvalidLotteryId
	}
	if lottery.BidCount > 10 {
		return sdkerrors.Wrapf(fmt.Errorf("lottery: %s", lottery.Index), ErrFinishedLottery.Error())
	}
	if lottery.MinBetAmount < 5 {
		return sdkerrors.Wrapf(fmt.Errorf("lottery: %s", lottery.Index), ErrInvalidBetAmount.Error())
	}

	if lottery.Fee < 1 {
		return sdkerrors.Wrapf(fmt.Errorf("lottery: %s", lottery.Index), ErrInvalidFeeAmount.Error())
	}
	return nil
}

func (bid StoredBid) ValidateBid() error {
	_, err := bid.GetBidderAddress()
	if err != nil {
		return err
	}
	id, err := GenValidBidIndex(bid.Bidder, bid.LotteryId)
	if err != nil {
		return err
	}
	if id != bid.Index {
		return sdkerrors.Wrapf(fmt.Errorf("bid: %s, lotteryId: %d", bid.Index, bid.LotteryId), ErrInvalidLotteryId.Error())
	}
	if bid.BetAmount < 5 {
		return ErrInvalidBetAmount
	}
	return nil
}
