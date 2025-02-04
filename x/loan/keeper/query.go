package keeper

import (
	"example/x/loan/types"
)

var _ types.QueryServer = Keeper{}
