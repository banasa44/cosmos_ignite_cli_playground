package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"example/x/blog/keeper"
	"example/x/blog/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)
	// Set all the sentPost
	for _, elem := range genState.SentPostList {
		k.SetSentPost(ctx, elem)
	}

	// Set sentPost count
	k.SetSentPostCount(ctx, genState.SentPostCount)
	// Set all the timeoutPost
	for _, elem := range genState.TimeoutPostList {
		k.SetTimeoutPost(ctx, elem)
	}

	// Set timeoutPost count
	k.SetTimeoutPostCount(ctx, genState.TimeoutPostCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.PostList = k.GetAllPost(ctx)
	genesis.PostCount = k.GetPostCount(ctx)
	genesis.SentPostList = k.GetAllSentPost(ctx)
	genesis.SentPostCount = k.GetSentPostCount(ctx)
	genesis.TimeoutPostList = k.GetAllTimeoutPost(ctx)
	genesis.TimeoutPostCount = k.GetTimeoutPostCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
