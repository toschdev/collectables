package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/tosch110/collectables/x/collectables"
)

// OverrideNFTModule overrides the NFT module for custom handlers
type OverrideNFTModule struct {
	nft.AppModule
	k nft.Keeper
}

// NewHandler module handler for the OerrideNFTModule
func (am OverrideNFTModule) NewHandler() sdk.Handler {
	return CustomNFTHandler(am.k)
}

// NewOverrideNFTModule generates a new NFT Module
func NewOverrideNFTModule(appModule nft.AppModule, keeper nft.Keeper) OverrideNFTModule {
	return OverrideNFTModule{
		AppModule: appModule,
		k:         keeper,
	}
}

// CustomNFTHandler routes the messages to the handlers
func CustomNFTHandler(k nft.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case nft.MsgSendNFT:
			return nft.HandleMsgSendNFT(ctx, msg, k)
		case nft.MsgEditNFTMetadata:
			return nft.HandleMsgEditNFTMetadata(ctx, msg, k)
		case nft.MsgMintNFT:
			return HandleMsgMintNFTCustom(ctx, msg, k)
		case nft.MsgBuyNFT:
			return nft.HandleMsgBuyNFT(ctx, msg, k)
		case nft.MsgSellNFT:
			return nft.HandleMsgSellNFT(ctx, msg, k)
		case nft.MsgChallengeNFT:
			return nft.HandleMsgChallengeNFTCustom(ctx, msg, k)
		case nft.MsgChallengeNFTProof:
			return nft.HandleMsgChallengeNFTProof(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized nft message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// HandleMsgChallengeNFTCustom handles HandleMsgChallengeNFT
func HandleMsgChallengeNFTCustom(ctx sdk.Context, msg nft.MsgChallengeNFT, k nft.Keeper,
) sdk.Result {

	// Check if owner of both tokens is not the same
	// Check who is going to win
	// Handle transaction accordingly
	isTwilight := checkTwilightChallenge(ctx)

	if isTwilight {
		return nft.HandleMsgChallengeNFT(ctx, msg, k)
	}

	errMsg := fmt.Sprintf("Can't challenge your own tokens.")
	return sdk.ErrUnknownRequest(errMsg).Result()
}

// HandleMsgMintNFTCustom handles MsgMintNFT
func HandleMsgMintNFTCustom(ctx sdk.Context, msg nft.MsgMintNFT, k nft.Keeper,
) sdk.Result {

	// Check for Blake3 Proof
	// Check for timestamp
	// Check if Wins and Losses are 0
	isTwilight := checkTwilight(ctx)

	if isTwilight {
		return nft.HandleMsgChallengeNFT(ctx, msg, k)
	}

	errMsg := fmt.Sprintf("Not a valid NFT Token!")
	return sdk.ErrUnknownRequest(errMsg).Result()
}

func checkTwilight(ctx sdk.Context) bool {
	header := ctx.BlockHeader()
	time := header.Time
	fmt.Println("time", time)
	return true
}

func checkTwilightChallenge(ctx sdk.Context) bool {
	header := ctx.BlockHeader()
	time := header.Time
	fmt.Println("time", time)
	return true
}
