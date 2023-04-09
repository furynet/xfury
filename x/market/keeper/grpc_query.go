package keeper

import (
	"github.com/furynet/xfury/x/market/types"
)

var _ types.QueryServer = Keeper{}
