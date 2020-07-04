package nft

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tosch110/collectables/x/collectables/keeper"
	"github.com/tosch110/collectables/x/collectables/types"
)

// GenericHandler routes the messages to the handlers
func GenericHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgSendNFT:
			return HandleMsgSendNFT(ctx, msg, k)
		case types.MsgEditNFTMetadata:
			return HandleMsgEditNFTMetadata(ctx, msg, k)
		case types.MsgEditNFTPrice:
			return HandleMsgEditNFTPrice(ctx, msg, k)
		case types.MsgMintNFT:
			return HandleMsgMintNFT(ctx, msg, k)
		case types.MsgBurnNFT:
			return HandleMsgBurnNFT(ctx, msg, k)
		case types.MsgBuyNFT:
			return HandleMsgBuyNFT(ctx, msg, k)
		case types.MsgChallengeNFT:
			return HandleMsgChallengeNFT(ctx, msg, k)
		case types.MsgChallengeNFTProof:
			return HandleMsgChallengeNFTProof(ctx, msg, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unrecognized nft message type: %T", msg))
		}
	}
}

// HandleMsgSendNFT handler for MsgSendNFT
func HandleMsgSendNFT(ctx sdk.Context, msg types.MsgSendNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}
	// update NFT owner
	nft.SetOwner(msg.Recipient)
	// update the NFT (owners are updated within the keeper)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgEditNFTMetadata handler for MsgEditNFTMetadata
func HandleMsgEditNFTMetadata(ctx sdk.Context, msg types.MsgEditNFTMetadata, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// update NFT
	nft.EditMetadata(msg.TokenURI)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFTMetadata,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTTokenURI, msg.TokenURI),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgEditNFTPrice handler for MsgEditNFTPrice
func HandleMsgEditNFTPrice(ctx sdk.Context, msg types.MsgEditNFTPrice, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// update NFT
	nft.EditPrice(msg.Price)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFTPrice,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTTokenURI, msg.Price),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgMintNFT handles MsgMintNFT
func HandleMsgMintNFT(ctx sdk.Context, msg types.MsgMintNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	nft := types.NewBaseNFT(msg.ID, msg.Recipient, msg.Hash, msg.Proof, msg.Name, 0, 0, 0)
	err := k.MintNFT(ctx, msg.Denom, &nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintNFT,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTHash, msg.Hash),
			sdk.NewAttribute(types.AttributeKeyNFTProof, msg.Proof),
			sdk.NewAttribute(types.AttributeKeyNFTName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyNFTWins, 0),
			sdk.NewAttribute(types.AttributeKeyNFTLosses, 0),
			sdk.NewAttribute(types.AttributeKeyNFTPrice, 0),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgBurnNFT handles MsgBurnNFT
func HandleMsgBurnNFT(ctx sdk.Context, msg types.MsgBurnNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	_, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// remove  NFT
	err = k.DeleteNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnNFT,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgBuyNFT handler for MsgBuyNFT
func HandleMsgBuyNFT(ctx sdk.Context, msg types.MsgBuyNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// Checks if the the buy price is equal to the price set by the current owner
	if nft.GetPrice().IsAllGT(msg.Price) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Not enough coins provided.") // If not, throw an error
	}

	// Price matches, send coins
	err := keeper.CoinKeeper.SendCoins(ctx, msg.Sender, nft.GetOwner(), msg.Price)
	if err != nil {
		return nil, err
	}

	// update NFT owner
	nft.SetOwner(msg.Sender)
	nft.SetPrice(0)
	// update the NFT (owners are updated within the keeper)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// EndBlocker is run at the end of the block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return nil
}
