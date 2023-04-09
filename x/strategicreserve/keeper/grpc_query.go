package keeper

import (
	"github.com/furynet/fury/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
