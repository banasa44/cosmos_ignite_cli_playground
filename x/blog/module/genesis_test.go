package blog_test

import (
	"testing"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	blog "example/x/blog/module"
	"example/x/blog/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		SentPostList: []types.SentPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentPostCount: 2,
		TimeoutPostList: []types.TimeoutPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimeoutPostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, k, genesisState)
	got := blog.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	require.ElementsMatch(t, genesisState.SentPostList, got.SentPostList)
	require.Equal(t, genesisState.SentPostCount, got.SentPostCount)
	require.ElementsMatch(t, genesisState.TimeoutPostList, got.TimeoutPostList)
	require.Equal(t, genesisState.TimeoutPostCount, got.TimeoutPostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
