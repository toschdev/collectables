package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nft "github.com/tosch110/collectables/x/collectables"
	"github.com/tosch110/collectables/x/collectables/types"
)

// OverrideNFTModule overrides the NFT module for custom handlers
type OverrideNFTModule struct {
	nft.AppModule
	k nft.Keeper
}

// NewHandler module handler for the OerrideNFTModule
func (am OverrideNFTModule) NewHandler() sdk.Handler {
	k := am.k
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case nft.MsgSendNFT:
			result, err := nft.HandleMsgSendNFT(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Sending not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		case nft.MsgEditNFTMetadata:
			result, err := nft.HandleMsgEditNFTMetadata(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Edit Metadata not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		case nft.MsgEditNFTPrice:
			result, err := nft.HandleMsgEditNFTPrice(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Edit Price not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		case nft.MsgMintNFT:
			result, err := nft.HandleMsgMintNFT(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Mint NFT not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		case nft.MsgBuyNFT:
			result, err := nft.HandleMsgBuyNFT(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Edit Price not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		case nft.MsgChallengeNFT:
			result, err := nft.HandleMsgChallengeNFT(ctx, msg, k)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
					fmt.Sprintf("Challenge NFT not successful %s : %T", types.ModuleName, msg))
			}
			return result, nil
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("Challenge NFT not successful %s : %T", types.ModuleName, msg))
		}
	}
}

// NewOverrideNFTModule generates a new NFT Module
func NewOverrideNFTModule(appModule nft.AppModule, keeper nft.Keeper) OverrideNFTModule {
	return OverrideNFTModule{
		AppModule: appModule,
		k:         keeper,
	}
}
