package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/params/msgstat"
)

type GenesisState struct {
	MsgStatusState msgstat.GenesisState
}

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	msgstat.InitGenesis(ctx, k.MsgStatusStore(), data.MsgStatusState)
}
