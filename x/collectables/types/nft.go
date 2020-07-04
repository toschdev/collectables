package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// var _ NFT = (*BaseNFT)(nil)

// BaseNFT non fungible token definition
type BaseNFT struct {
	ID    string         `json:"id,omitempty" yaml:"id"` // id of the token; not exported to clients
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`     // account address that owns the NFT
	//gamification data
	Hash   string    `json:"hash" yaml:"hash"`     // blake3 hash of the proof
	Proof  string    `json:"proof" yaml:"proof"`   // Proof, input for the hash (25-255 character)
	Name   string    `json:"name" yaml:"name"`     // The user given name of the NFT Token
	Wins   uint      `json:"wins" yaml:"wins"`     // Challenge Wins
	Losses uint      `json:"losses" yaml:"losses"` // Challenge Losses
	Price  sdk.Coins `json:"price" yaml:"price"`   // The price set for the token
}

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id string, owner sdk.AccAddress, hash, proof, name string, wins, losses uint, price sdk.Coins) BaseNFT {
	return BaseNFT{
		ID:     id,
		Owner:  owner,
		Hash:   strings.TrimSpace(hash),
		Proof:  strings.TrimSpace(proof),
		Name:   strings.TrimSpace(name),
		Wins:   wins,
		Losses: losses,
		Price:  price,
	}
}

// GetID returns the ID of the token
func (bnft BaseNFT) GetID() string { return bnft.ID }

// GetOwner returns the account address that owns the NFT
func (bnft BaseNFT) GetOwner() sdk.AccAddress { return bnft.Owner }

// SetOwner updates the owner address of the NFT
func (bnft *BaseNFT) SetOwner(address sdk.AccAddress) {
	bnft.Owner = address
}

// GetHash returns the path to optional extra properties
func (bnft BaseNFT) GetHash() string { return bnft.Hash }

// GetProof returns the path to optional extra properties
func (bnft BaseNFT) GetProof() string { return bnft.Proof }

// GetName returns the path to optional extra properties
func (bnft BaseNFT) GetName() string { return bnft.Name }

// GetWins returns the path to optional extra properties
func (bnft BaseNFT) GetWins() uint { return bnft.Wins }

// GetLosses returns the path to optional extra properties
func (bnft BaseNFT) GetLosses() uint { return bnft.Losses }

// GetPrice returns the price for an NFT Token
func (bnft *BaseNFT) GetPrice() sdk.Coins { return bnft.Price }

// EditPrice removes an Ask order to an nft.
func (bnft *BaseNFT) EditPrice(price sdk.Coins) {
	bnft.Price = price
}

// EditMetadata edits metadata of an nft
func (bnft *BaseNFT) EditMetadata(name string) {
	bnft.Name = name
}

// IncreaseWins inceases wins of an nft
func (bnft *BaseNFT) IncreaseWins() {
	bnft.Wins = bnft.Wins + 1
}

// IncreaseLosses inceases wins of an nft
func (bnft *BaseNFT) IncreaseLosses() {
	bnft.Wins = bnft.Losses + 1
}

func (bnft BaseNFT) String() string {
	return fmt.Sprintf(`ID:				%s
Owner:			%s
Hash:		%s
Proof:		%s
Name:		%s
Wins:       %v
Losses:     %v
Price:      %v`,
		bnft.ID,
		bnft.Owner,
		bnft.Hash,
		bnft.Proof,
		bnft.Name,
		bnft.Wins,
		bnft.Losses,
		bnft.Price,
	)
}

// ----------------------------------------------------------------------------
// NFT

// NFTs define a list of NFT
type NFTs []NFT

// NewNFTs creates a new set of NFTs
func NewNFTs(nfts ...NFT) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return NFTs(nfts).Sort()
}

// Append appends two sets of NFTs
func (nfts NFTs) Append(nftsB ...NFT) NFTs {
	return append(nfts, nftsB...).Sort()
}

// Find returns the searched collection from the set
func (nfts NFTs) Find(id string) (nft NFT, found bool) {
	index := nfts.find(id)
	if index == -1 {
		return nft, false
	}
	return nfts[index], true
}

// Update removes and replaces an NFT from the set
func (nfts NFTs) Update(id string, nft NFT) (NFTs, bool) {
	index := nfts.find(id)
	if index == -1 {
		return nfts, false
	}

	return append(append(nfts[:index], nft), nfts[index+1:]...), true
}

// Remove removes an NFT from the set of NFTs
func (nfts NFTs) Remove(id string) (NFTs, bool) {
	index := nfts.find(id)
	if index == -1 {
		return nfts, false
	}

	return append(nfts[:index], nfts[index+1:]...), true
}

// String follows stringer interface
func (nfts NFTs) String() string {
	if len(nfts) == 0 {
		return ""
	}

	out := ""
	for _, nft := range nfts {
		out += fmt.Sprintf("%v\n", nft.String())
	}
	return out[:len(out)-1]
}

// Empty returns true if there are no NFTs and false otherwise.
func (nfts NFTs) Empty() bool {
	return len(nfts) == 0
}

func (nfts NFTs) find(id string) int {
	return FindUtil(nfts, id)
}

// ----------------------------------------------------------------------------
// Encoding

// NFTJSON is the exported NFT format for clients
type NFTJSON map[string]BaseNFT

// MarshalJSON for NFTs
func (nfts NFTs) MarshalJSON() ([]byte, error) {
	nftJSON := make(NFTJSON)
	for _, nft := range nfts {
		id := nft.GetID()
		bnft := NewBaseNFT(id, nft.GetOwner(), nft.GetHash(), nft.GetProof(), nft.GetName(), nft.GetWins(), nft.GetLosses(), nft.GetPrice())
		nftJSON[id] = bnft
	}
	return json.Marshal(nftJSON)
}

// UnmarshalJSON for NFTs
func (nfts *NFTs) UnmarshalJSON(b []byte) error {
	nftJSON := make(NFTJSON)
	if err := json.Unmarshal(b, &nftJSON); err != nil {
		return err
	}

	for id, nft := range nftJSON {
		bnft := NewBaseNFT(id, nft.GetOwner(), nft.GetHash(), nft.GetProof(), nft.GetName(), nft.GetWins(), nft.GetLosses(), nft.GetPrice())
		*nfts = append(*nfts, &bnft)
	}
	return nil
}

// Findable and Sort interfaces
func (nfts NFTs) ElAtIndex(index int) string { return nfts[index].GetID() }
func (nfts NFTs) Len() int                   { return len(nfts) }
func (nfts NFTs) Less(i, j int) bool         { return strings.Compare(nfts[i].GetID(), nfts[j].GetID()) == -1 }
func (nfts NFTs) Swap(i, j int)              { nfts[i], nfts[j] = nfts[j], nfts[i] }

var _ sort.Interface = NFTs{}

// Sort is a helper function to sort the set of coins in place
func (nfts NFTs) Sort() NFTs {
	sort.Sort(nfts)
	return nfts
}
