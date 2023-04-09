package keeper

import (
	"github.com/furynet/fury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
