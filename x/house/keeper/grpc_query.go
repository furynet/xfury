package keeper

import (
	"github.com/furynet/fury/x/house/types"
)

var _ types.QueryServer = Keeper{}
