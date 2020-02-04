package bep3

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/bep3/internal/types"
)

// NewHandler creates an sdk.Handler for all the bep3 type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateHTLT:
			return handleMsgCreateHTLT(ctx, k, msg)
		// TODO: Define your msg cases
		//
		//Example:
		// case MsgSet<Action>:
		// 	return handleMsg<Action>(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgCreateHTLT handles requests to create a new HTLT
func handleMsgCreateHTLT(ctx sdk.Context, k Keeper, msg types.MsgCreateHTLT) sdk.Result {
	// msg contains HTLT attributes
	// Validate Chain name
	// if name == "kava", initiator must be relayer

	err := k.AddHTLT(ctx, msg.From, msg.To, msg.RecipientOtherChain,
		msg.SenderOtherChain, msg.RandomNumberHash, msg.Timestamp, msg.Amount,
		msg.ExpectedIncome, msg.HeightSpan, msg.CrossChain)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)
	id, _ := k.GetHtltID(ctx, msg.Sender, msg.Collateral[0].Denom)

	return sdk.Result{
		Data:   GetHtltIDBytes(id),
		Events: ctx.EventManager().Events(),
	}
}