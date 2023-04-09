package keeper

import (
	"github.com/furynet/fury/x/mint/types"
)

var _ types.QueryServer = Keeper{}
