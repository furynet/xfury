package keeper

import (
	"github.com/furynet/xfury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
