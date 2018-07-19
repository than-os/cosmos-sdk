package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/params/gas"
	"github.com/cosmos/cosmos-sdk/x/params/msgstat"
)

func NewAnteHandler(k Keeper) sdk.AnteHandler {
	gasante := gas.NewAnteHandler(k.GasConfigStore())
	msgstatante := msgstat.NewAnteHandler(k.MsgStatusStore())

	return func(ctx sdk.Context, tx sdk.Tx) (newctx sdk.Context, res sdk.Result, abort bool) {
		newctx, res, abort = gasante(ctx, tx)
		if abort {
			return
		}
		return msgstatante(newctx, tx)
	}
}
