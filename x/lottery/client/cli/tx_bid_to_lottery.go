package cli

import (
	"strconv"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBidToLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bid-to-lottery [lottery-id] [bid-amount] [bid-count]",
		Short: "Broadcast message bidToLottery",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLotteryId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argBidAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argBidCount, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBidToLottery(
				clientCtx.GetFromAddress().String(),
				argLotteryId,
				argBidAmount,
				argBidCount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
