package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* --------------------------------------------------------------------------- */
// MsgSendNFT
/* --------------------------------------------------------------------------- */

// MsgSendNFT defines a TransferNFT message
type MsgSendNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Denom     string         `json:"denom" yaml:"denom"`
	ID        string         `json:"id" yaml:"id"`
}

// NewMsgSendNFT is a constructor function for MsgSetName
func NewMsgSendNFT(sender, recipient sdk.AccAddress, denom, id string) MsgSendNFT {
	return MsgSendNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
	}
}

// Route Implements Msg
func (msg MsgSendNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgSendNFT) Type() string { return "send_nft" }

// ValidateBasic Implements Msg.
func (msg MsgSendNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidCollection
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSendNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSendNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgEditNFTMetadata
/* --------------------------------------------------------------------------- */

// MsgEditNFTMetadata edits an NFT's metadata
type MsgEditNFTMetadata struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	ID     string         `json:"id" yaml:"id"`
	Denom  string         `json:"denom" yaml:"denom"`
	Name   string         `json:"name" yaml:"name"`
}

// NewMsgEditNFTMetadata is a constructor function for MsgSetName
func NewMsgEditNFTMetadata(sender sdk.AccAddress, id,
	denom, name string,
) MsgEditNFTMetadata {
	return MsgEditNFTMetadata{
		Sender: sender,
		ID:     strings.TrimSpace(id),
		Denom:  strings.TrimSpace(denom),
		Name:   strings.TrimSpace(name),
	}
}

// Route Implements Msg
func (msg MsgEditNFTMetadata) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFTMetadata) Type() string { return "edit_nft_metadata" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFTMetadata) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFTMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditNFTMetadata) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgMintNFT
/* --------------------------------------------------------------------------- */

// MsgMintNFT defines a MintNFT message
type MsgMintNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	Hash      string         `json:"hash" yaml:"hash"`
	Proof     string         `json:"proof" yaml:"proof"`
	Name      string         `json:"name" yaml:"name"`
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id, denom, hash, proof, name string) MsgMintNFT {
	return MsgMintNFT{
		Sender:    sender,
		Recipient: recipient,
		ID:        strings.TrimSpace(id),
		Denom:     strings.TrimSpace(denom),
		Hash:      strings.TrimSpace(hash),
		Proof:     strings.TrimSpace(proof),
		Name:      strings.TrimSpace(name),
	}
}

// Route Implements Msg
func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "mint_nft" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgBurnNFT
/* --------------------------------------------------------------------------- */

// MsgBurnNFT defines a BurnNFT message
type MsgBurnNFT struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	ID     string         `json:"id" yaml:"id"`
	Denom  string         `json:"denom" yaml:"denom"`
}

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(sender sdk.AccAddress, id string, denom string) MsgBurnNFT {
	return MsgBurnNFT{
		Sender: sender,
		ID:     strings.TrimSpace(id),
		Denom:  strings.TrimSpace(denom),
	}
}

// Route Implements Msg
func (msg MsgBurnNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return "burn_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgBuyNFT
/* --------------------------------------------------------------------------- */

// MsgBuyNFT defines a BuyNFT message
type MsgBuyNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	Bid       sdk.Coins      `json:"bid" yaml:"bid"`
}

// NewMsgBuyNFT is a constructor function for MsgBuyNFT
func NewMsgBuyNFT(sender, recipient sdk.AccAddress, id, denom string, bid sdk.Coins) MsgBuyNFT {
	return NewMsgBuyNFT{
		Sender:    sender,
		Recipient: recipient,
		ID:        strings.TrimSpace(id),
		Denom:     strings.TrimSpace(denom),
		Bid:       bid,
	}
}

// Route Implements Msg
func (msg MsgBuyNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBuyNFT) Type() string { return "buy_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBuyNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBuyNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBuyNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgSellNFT
/* --------------------------------------------------------------------------- */

// MsgSellNFT defines a SellNFT message
type MsgSellNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	Ask       sdk.Coins      `json:"ask" yaml:"ask"`
}

// NewMsgSellNFT is a constructor function for MsgSellNFT
func NewMsgSellNFT(sender, recipient sdk.AccAddress, id, denom string, ask sdk.Coins) MsgSellNFT {
	return NewMsgSellNFT{
		Sender:    sender,
		Recipient: recipient,
		ID:        strings.TrimSpace(id),
		Denom:     strings.TrimSpace(denom),
		Ask:       ask,
	}
}

// Route Implements Msg
func (msg MsgSellNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgSellNFT) Type() string { return "sell_nft" }

// ValidateBasic Implements Msg.
func (msg MsgSellNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	if !msg.Ask.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSellNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSellNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgChallengeNFT
/* --------------------------------------------------------------------------- */

// MsgChallengeNFT defines a ChallengeNFT message
type MsgChallengeNFT struct {
	ContenderID    string `json:"contenderid" yaml:"contenderid"`
	ContenderDenom string `json:"contenderdenom" yaml:"contenderdenom"`
	DefiantID      string `json:"defiantid" yaml:"defiantid"`
	DefiantDenom   string `json:"defiantdenom" yaml:"defiantdenom"`
}

// NewMsgChallengeNFT is a constructor function for MsgChallengeNFT
func NewMsgChallengeNFT(contenderid, contenderdenom, defiantid, defiantdenom string) MsgChallengeNFT {
	return NewMsgChallengeNFT{
		ContenderDenom: strings.TrimSpace(contenderdenom),
		ContenderID:    strings.TrimSpace(contenderid),
		DefiantDenom:   strings.TrimSpace(defiantdenom)
		DefiantID:      strings.TrimSpace(defiantid),
	}
}

// Route Implements Msg
func (msg MsgSellNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgSellNFT) Type() string { return "challenge_nft" }

// ValidateBasic Implements Msg.
func (msg MsgSellNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.ContenderDenom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ContenderID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.DefiantDenom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.DefiantID) == "" {
		return ErrInvalidNFT
	}
	if msg.Contender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid contender address")
	}
	if msg.Defiant.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid defiant address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSellNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSellNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

/* --------------------------------------------------------------------------- */
// MsgChallengeNFTProof
/* --------------------------------------------------------------------------- */

// MsgChallengeNFTProof defines a ChallengeNFTProof message
type MsgChallengeNFTProof struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
}

// NewMsgChallengeNFTProofis a constructor function for MsgChallengeNFTProof
func NewMsgChallengeNFTProof(sender, recipient sdk.AccAddress, id, denom, hash, proof string) MsgChallengeNFTProof {
	return MsgChallengeNFTProof{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
	}
}

// Route Implements Msg
func (msg MsgChallengeNFTProof) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgChallengeNFTProof) Type() string { return "challenge_nft_proof" }

// ValidateBasic Implements Msg.
func (msg MsgChallengeNFTProof) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgChallengeNFTProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgChallengeNFTProof) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
