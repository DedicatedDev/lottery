package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBidToLottery = "bid_to_lottery"

var _ sdk.Msg = &MsgBidToLottery{}

func NewMsgBidToLottery(creator string, lotteryId uint64, bidAmount uint64) *MsgBidToLottery {
	return &MsgBidToLottery{
		Creator:   creator,
		LotteryId: lotteryId,
		BidAmount: bidAmount,
	}
}

func (msg *MsgBidToLottery) Route() string {
	return RouterKey
}

func (msg *MsgBidToLottery) Type() string {
	return TypeMsgBidToLottery
}

func (msg *MsgBidToLottery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBidToLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBidToLottery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
