package keeper

import (
	"context"

	"example/x/example/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authority returns the module's authority address.
func (k Keeper) Authority(goCtx context.Context, req *types.QueryAuthorityRequest) (*types.QueryAuthorityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryAuthorityResponse{
		Address: k.GetAuthority(),
	}, nil
}
