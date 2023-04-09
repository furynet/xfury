package keeper

import (
	"github.com/furynet/fury/x/market/types"
)

var _ types.QueryServer = Keeper{}
