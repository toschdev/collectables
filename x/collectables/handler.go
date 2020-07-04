package collectables

import (
	"encoding/hex"
	"fmt"
	"lukechampine.com/blake3"
	"strconv"

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
			types.EventTypeSend,
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
	nft.EditMetadata(msg.Name)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFTMetadata,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTName, msg.Name),
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
			sdk.NewAttribute(types.AttributeKeyNFTPrice, msg.Price.String()),
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
	nft := types.NewBaseNFT(msg.ID, msg.Recipient, msg.Hash, msg.Proof, msg.Name, 0, 0, msg.Price)
	err := k.MintNFT(ctx, msg.Denom, &nft)
	if err != nil {
		return nil, err
	}

	proofCheck := blakeHash(msg.Proof)
	if msg.Hash != proofCheck {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Hash is not the blake3 hash of the Proof") // If not, throw an error
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
			sdk.NewAttribute(types.AttributeKeyNFTWins, "0"),
			sdk.NewAttribute(types.AttributeKeyNFTLosses, "0"),
			sdk.NewAttribute(types.AttributeKeyNFTPrice, msg.Price.String()),
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

	// Checks if NFT is for sale
	if nft.GetPrice().IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Token is not for sale.") // If not, throw an error
	}

	// Checks if the the buy price is equal to the price set by the current owner
	if nft.GetPrice().IsAllGT(msg.Price) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Not enough coins provided.") // If not, throw an error
	}

	// Price matches, send coins
	err = k.CoinKeeper.SendCoins(ctx, msg.Sender, nft.GetOwner(), msg.Price)
	if err != nil {
		return nil, err
	}

	// update NFT owner
	nft.SetOwner(msg.Sender)
	// update the NFT (owners are updated within the keeper)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSend,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender.String()),
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

// HandleMsgChallengeNFT handler for MsgChallengeNFT
func HandleMsgChallengeNFT(ctx sdk.Context, msg types.MsgChallengeNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	contenderNFT, err := k.GetNFT(ctx, msg.ContenderDenom, msg.ContenderID)
	if err != nil {
		return nil, err
	}

	defiantNFT, err := k.GetNFT(ctx, msg.DefiantDenom, msg.DefiantID)
	if err != nil {
		return nil, err
	}

	contendantHash := contenderNFT.GetHash()
	contendantWins := contenderNFT.GetWins()

	defiantHash := defiantNFT.GetHash()
	defiantWins := defiantNFT.GetWins()

	matchResults, err := fight(contendantHash, int(contendantWins), defiantHash, int(defiantWins)) // our match logic. Returns the results for both tokens and the winner
	if err != nil {
		return nil, err
	}

	if matchResults.Winner == "contestant" {
		// update NFT owner
		defiantNFT.SetOwner(msg.Sender)
		// increase loose streak of looser of this match
		defiantNFT.IncreaseLosses()
		// increase winning streak of winner of this match
		contenderNFT.IncreaseWins()
	}

	if matchResults.Winner == "defiant" {
		// increase loose streak of looser of this match
		contenderNFT.IncreaseLosses()
		// increase winning streak of winner of this match
		defiantNFT.IncreaseWins()
	}

	// update the NFTs
	err = k.UpdateNFT(ctx, msg.ContenderDenom, contenderNFT)
	if err != nil {
		return nil, err
	}

	err = k.UpdateNFT(ctx, msg.DefiantDenom, defiantNFT)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChallengeNFT,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.ContenderDenom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ContenderID),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.DefiantDenom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.DefiantID),
			sdk.NewAttribute(types.AttributeKeyNFTWinner, msg.Winner),
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

func blakeHash(proof string) string {
	hash := blake3.New(32, nil)
	hash.Write([]byte(proof))
	output := toHex(hash.Sum(nil))

	return output
}

func toHex(data []byte) string { return hex.EncodeToString(data) }

// Fight struct for the fight information
type Fight struct {
	ContestantSum  int    `json:"contestant_sum"`
	ContestantHash string `json:"contestant_hash"`
	DefiantSum     int    `json:"defiant_sum"`
	DefiantHash    string `json:"defiant_hash"`
	Winner         string `json:"winner"`
}

func fight(contestantHash string, contestantWins int, defiantHash string, defiantWins int) (Fight, error) {

	var contestantSum int
	var defiantSum int

	contestantSum = contestantWins
	defiantSum = defiantWins

	var result Fight
	for _, s := range contestantHash {

		charValue := fmt.Sprintf("%d", s)

		runeValue, err := strconv.Atoi(charValue)
		if err != nil {
			return result, err
		}

		contestantSum = contestantSum + runeValue
	}

	for _, s := range defiantHash {

		charValue := fmt.Sprintf("%d", s)

		runeValue, err := strconv.Atoi(charValue)
		if err != nil {
			return result, err
		}

		defiantSum = defiantSum + runeValue
	}

	contestorWins := contestantSum - defiantSum

	winner := "defiant"
	if contestorWins > 0 {
		winner = "contestor"
	}

	result = Fight{
		ContestantHash: contestantHash,
		ContestantSum:  contestantSum,
		DefiantHash:    defiantHash,
		DefiantSum:     defiantSum,
		Winner:         winner,
	}

	return result, nil
}
