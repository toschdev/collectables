package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// The NFT interface
type NFT interface {
	GetID() string
	GetOwner() sdk.AccAddress
	SetOwner(address sdk.AccAddress)
	//the following functions are for our gamification module
	GetHash() string
	GetProof() string
	GetName() string
	GetWins() uint
	GetLosses() uint
	EditMetadata(tokenName string)
	String() string
}
