package types

// NFT module event types
var (
	EventTypeSend              = "send_nft"
	EventTypeEditNFTMetadata   = "edit_nft_metadata"
	EventTypeMintNFT           = "mint_nft"
	EventTypeBuyNFT            = "buy_nft"
	EventTypeSellNFT           = "sell_nft"
	EventTypeBurnNFT           = "burn_nft"
	EventTypeChallengeNFT      = "challenge_nft"
	EventTypeChallengeNFTProof = "challenge_nft_proof"

	AttributeValueCategory = ModuleName

	AttributeKeySender    = "sender"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyNFTID     = "nft-id"
	AttributeKeyNFTName   = "name"
	AttributeKeyNFTHash   = "hash"
	AttributeKeyNFTProof  = "proof"
	AttributeKeyWins      = "wins"
	AttributeKeyLosses    = "losses"
	AttributeKeyDenom     = "denom"
)
