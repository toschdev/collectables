package types

// NFT module event types
var (
	EventTypeSend            = "send_nft"
	EventTypeEditNFTMetadata = "edit_nft_metadata"
	EventTypeMintNFT         = "mint_nft"
	EventTypeBuyNFT          = "buy_nft"
	EventTypeEditNFTPrice    = "edit_nft_price"
	EventTypeBurnNFT         = "burn_nft"
	EventTypeChallengeNFT    = "challenge_nft"

	AttributeValueCategory = ModuleName

	AttributeKeySender    = "sender"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyNFTID     = "nft-id"
	AttributeKeyNFTName   = "name"
	AttributeKeyNFTHash   = "hash"
	AttributeKeyNFTProof  = "proof"
	AttributeKeyNFTWins   = "wins"
	AttributeKeyNFTLosses = "losses"
	AttributeKeyDenom     = "denom"
	AttributeKeyNFTPrice  = "price"
	AttributeKeyNFTWinner = "winner"
)
