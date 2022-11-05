package types_test

import (
	"fmt"
	"testing"

	//"fmt"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	//"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

// generate mock stored game data
func getStoredEmptyLottery() *types.StoredLottery {
	return &types.StoredLottery{
		Index:        "",
		BidCount:     0,
		MinBetAmount: 5,
		Fee:          1,
	}
}
func getStoredLottery() *types.StoredLottery {
	return &types.StoredLottery{
		Index:        "1",
		BidCount:     0,
		MinBetAmount: 5,
		Fee:          1,
	}
}
func getStoredLotteryWithInvalidMinBetAmount() *types.StoredLottery {
	return &types.StoredLottery{
		Index:        "1",
		BidCount:     0,
		MinBetAmount: 1,
		Fee:          1,
	}
}

func getStoredLotteryWithInvalidMinFee() *types.StoredLottery {
	return &types.StoredLottery{
		Index:        "1",
		BidCount:     0,
		MinBetAmount: 5,
		Fee:          0,
	}
}

// generate moc bid data
func getStoredBid() *types.StoredBid {
	index, _ := types.GenValidBidIndex(alice, 1)
	return &types.StoredBid{
		Index:     index,
		LotteryId: 1,
		BetAmount: 6,
		Bidder:    alice,
	}
}

func getStoredBidWithInvalidId() *types.StoredBid {
	return &types.StoredBid{
		Index:     "1",
		LotteryId: 1,
		BetAmount: 6,
		Bidder:    alice,
	}
}

func getStoredBidWithInvalidBetAmount() *types.StoredBid {
	index, _ := types.GenValidBidIndex(alice, 1)
	return &types.StoredBid{
		Index:     index,
		LotteryId: 1,
		BetAmount: 2,
		Bidder:    alice,
	}
}

func getStoredBidWithInvalidBetAddress() *types.StoredBid {
	index, _ := types.GenValidBidIndex(alice, 1)
	return &types.StoredBid{
		Index:     index,
		LotteryId: 1,
		BetAmount: 2,
		Bidder:    "",
	}
}
func TestValidateLottery(t *testing.T) {
	//test lottery index is valid or not.
	storedLottery := getStoredEmptyLottery()
	err := storedLottery.ValidateLottery()
	require.EqualValues(t, err, types.ErrInvalidLotteryId)

	//test lottery bid amount adopt the problem's requirement
	storedLottery = getStoredLotteryWithInvalidMinBetAmount()
	err = storedLottery.ValidateLottery()
	expectedErr := sdkerrors.Wrapf(fmt.Errorf("lottery: %s", storedLottery.Index), types.ErrInvalidBetAmount.Error()).Error()
	require.Equal(t, err.Error(), expectedErr)

	//test lottery fee amount adopt the problem's requirement
	storedLottery = getStoredLotteryWithInvalidMinFee()
	err = storedLottery.ValidateLottery()
	expectedErr = sdkerrors.Wrapf(fmt.Errorf("lottery: %s", storedLottery.Index), types.ErrInvalidFeeAmount.Error()).Error()
	require.Equal(t, err.Error(), expectedErr)

	//pass valid lottery
	storedLottery = getStoredLottery()
	err = storedLottery.ValidateLottery()
	require.Nil(t, err)
}

func TestValidateBid(t *testing.T) {
	//test bid index is valid or not.
	storedBid := getStoredBidWithInvalidId()
	err := storedBid.ValidateBid()
	expectedErr := sdkerrors.Wrapf(fmt.Errorf("bid: %s, lotteryId: %d", storedBid.Index, storedBid.LotteryId), types.ErrInvalidLotteryId.Error()).Error()
	require.EqualValues(t, err.Error(), expectedErr)

	//test bid amount adopt the problem's requirement
	storedBid = getStoredBidWithInvalidBetAddress()
	addressErr := storedBid.ValidateBid()
	_, err = sdk.AccAddressFromBech32(storedBid.Bidder)
	require.Equal(t, addressErr, err)

	//test bid amount adopt the problem's requirement
	storedBid = getStoredBidWithInvalidBetAmount()
	err = storedBid.ValidateBid()
	require.Equal(t, err, types.ErrInvalidBetAmount)
}
