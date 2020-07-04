package collectables

import (
	"github.com/tosch110/collectables/x/collectables/keeper"
	"github.com/tosch110/collectables/x/collectables/types"
)

const (
	QuerySupply       = keeper.QuerySupply
	QueryOwner        = keeper.QueryOwner
	QueryOwnerByDenom = keeper.QueryOwnerByDenom
	QueryCollection   = keeper.QueryCollection
	QueryDenoms       = keeper.QueryDenoms
	QueryNFT          = keeper.QueryNFT
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
)

var (
	// functions aliases
	RegisterInvariants       = keeper.RegisterInvariants
	AllInvariants            = keeper.AllInvariants
	SupplyInvariant          = keeper.SupplyInvariant
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RegisterCodec            = types.RegisterCodec
	NewCollection            = types.NewCollection
	EmptyCollection          = types.EmptyCollection
	NewCollections           = types.NewCollections
	ErrInvalidCollection     = types.ErrInvalidCollection
	ErrUnknownCollection     = types.ErrUnknownCollection
	ErrInvalidNFT            = types.ErrInvalidNFT
	ErrNFTAlreadyExists      = types.ErrNFTAlreadyExists
	ErrUnknownNFT            = types.ErrUnknownNFT
	ErrEmptyMetadata         = types.ErrEmptyMetadata
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	GetCollectionKey         = types.GetCollectionKey
	SplitOwnerKey            = types.SplitOwnerKey
	GetOwnersKey             = types.GetOwnersKey
	GetOwnerKey              = types.GetOwnerKey
	NewMsgSendNFT            = types.NewMsgSendNFT
	NewMsgEditNFTMetadata    = types.NewMsgEditNFTMetadata
	NewMsgEditNFTPrice       = types.NewMsgEditNFTPrice
	NewMsgMintNFT            = types.NewMsgMintNFT
	NewMsgBurnNFT            = types.NewMsgBurnNFT
	NewMsgBuyNFT             = types.NewMsgBuyNFT
	NewMsgChallengeNFT       = types.NewMsgChallengeNFT
	NewMsgChallengeNFTProof  = types.NewMsgChallengeNFTProof
	NewBaseNFT               = types.NewBaseNFT
	NewNFTs                  = types.NewNFTs
	NewIDCollection          = types.NewIDCollection
	NewOwner                 = types.NewOwner
	NewQueryCollectionParams = types.NewQueryCollectionParams
	NewQueryBalanceParams    = types.NewQueryBalanceParams
	NewQueryNFTParams        = types.NewQueryNFTParams

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	EventTypeTransfer        = types.EventTypeTransfer
	EventTypeEditNFTMetadata = types.EventTypeEditNFTMetadata
	EventTypeMintNFT         = types.EventTypeMintNFT
	EventTypeBurnNFT         = types.EventTypeBurnNFT
	AttributeValueCategory   = types.AttributeValueCategory
	AttributeKeySender       = types.AttributeKeySender
	AttributeKeyRecipient    = types.AttributeKeyRecipient
	AttributeKeyOwner        = types.AttributeKeyOwner
	AttributeKeyNFTID        = types.AttributeKeyNFTID
	AttributeKeyNFTName      = types.AttributeKeyNFTName
	AttributeKeyNFTHash      = types.AttributeKeyNFTHash
	AttributeKeyNFTProof     = types.AttributeKeyNFTProof
	AttributeKeyNFTPrice     = types.AttributeKeyNFTPrice
	AttributeKeyDenom        = types.AttributeKeyDenom
	CollectionsKeyPrefix     = types.CollectionsKeyPrefix
	OwnersKeyPrefix          = types.OwnersKeyPrefix
)

type (
	Keeper                = keeper.Keeper
	Collection            = types.Collection
	Collections           = types.Collections
	CollectionJSON        = types.CollectionJSON
	GenesisState          = types.GenesisState
	MsgSendNFT            = types.MsgSendNFT
	MsgEditNFTMetadata    = types.MsgEditNFTMetadata
	MsgEditNFTPrice       = types.MsgEditNFTPrice
	MsgMintNFT            = types.MsgMintNFT
	MsgBurnNFT            = types.MsgBurnNFT
	MsgBuyNFT             = types.MsgBuyNFT
	MsgChallengeNFT       = types.MsgChallengeNFT
	MsgChallengeNFTProof  = types.MsgChallengeNFTProof
	BaseNFT               = types.BaseNFT
	NFTs                  = types.NFTs
	NFTJSON               = types.NFTJSON
	IDCollection          = types.IDCollection
	IDCollections         = types.IDCollections
	Owner                 = types.Owner
	QueryCollectionParams = types.QueryCollectionParams
	QueryBalanceParams    = types.QueryBalanceParams
	QueryNFTParams        = types.QueryNFTParams
)
