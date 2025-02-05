package keeper

import (
	"example/x/blog/types"
)

var _ types.QueryServer = Keeper{}
