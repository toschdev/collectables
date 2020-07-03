package types

// The NFT interface
type NFT interface {
	GetID() string
	GetOwner() sdk.AccAddress
	SetOwner(address sdk.AccAddress)
	//the following functions are for our gamification module
	GetTokenHash() string
	GetTokenProof() string
	GetTokenName() string
	GetWins() int
	GetLosses() int
	EditMetadata(tokenName string)
}
