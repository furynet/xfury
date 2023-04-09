package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/furynet/xfury/consts"
	"github.com/furynet/xfury/x/house/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Deposits queries all deposits
func (k Keeper) Deposits(c context.Context, req *types.QueryDepositsRequest) (*types.QueryDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest)
	}

	var deposits []types.Deposit
	ctx := sdk.UnwrapSDKContext(c)

	store := k.getDepositStore(ctx)
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var deposit types.Deposit
		if err := k.cdc.Unmarshal(value, &deposit); err != nil {
			return err
		}

		deposits = append(deposits, deposit)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDepositsResponse{Deposits: deposits, Pagination: pageRes}, nil
}

// DepositsByAccount returns all deposits of a given account address
func (k Keeper) DepositsByAccount(c context.Context,
	req *types.QueryDepositsByAccountRequest,
) (*types.QueryDepositsByAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest)
	}

	var deposits []types.Deposit
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(k.getDepositStore(ctx), types.GetDepositListPrefix(req.Address))
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var deposit types.Deposit
		if err := k.cdc.Unmarshal(value, &deposit); err != nil {
			return err
		}

		deposits = append(deposits, deposit)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDepositsByAccountResponse{Deposits: deposits, Pagination: pageRes}, nil
}
