package keeper

import (
	"github.com/mkXultra/oracle/x/consuming/types"
)

var _ types.QueryServer = Keeper{}
