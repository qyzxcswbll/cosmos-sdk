package fee_grant

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgGrantFeeAllowance:
			grant := FeeAllowanceGrant(msg)
			k.GrantFeeAllowance(ctx, grant)
			return sdk.Result{}
		case MsgRevokeFeeAllowance:
			k.RevokeFeeAllowance(ctx, msg.Granter, msg.Grantee)
			return sdk.Result{}
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %s", ModuleName)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
