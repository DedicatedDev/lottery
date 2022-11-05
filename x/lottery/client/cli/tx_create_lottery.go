package cli

import (
	"strconv"

	"github.com/DedicatedDev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-lottery [mint-bet-amount] [fee]",
		Short: "Broadcast message createLottery",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMintBetAmount, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			argFee, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLottery(
				clientCtx.GetFromAddress().String(),
				uint64(argMintBetAmount),
				uint64(argFee),
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
