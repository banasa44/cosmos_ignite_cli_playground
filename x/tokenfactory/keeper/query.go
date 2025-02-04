package keeper

import (
	"example/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
