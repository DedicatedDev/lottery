package keeper

import (
	"github.com/DedicatedDev/lottery/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
