package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/tosch110/collectables/x/collectables/types"

	"github.com/gorilla/mux"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router,
	cdc *codec.Codec, queryRoute string) {
	// Send an NFT to an address
	r.HandleFunc(
		"/nfts/send",
		sendNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Update an NFT metadata
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/metadata",
		editNFTMetadataHandler(cdc, cliCtx),
	).Methods("PUT")

	// Mint an NFT
	r.HandleFunc(
		"/nfts/mint",
		mintNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Burn an NFT
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/burn",
		burnNFTHandler(cdc, cliCtx),
	).Methods("PUT")

	// Challenge an NFT
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/challenge",
		challengeNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Buy an NFT
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/buy",
		buyNFTHandler(cdc, cliCtx),
	).Methods("POST")

	// Sell an NFT
	r.HandleFunc(
		"/nfts/collection/{denom}/nft/{id}/sell",
		sellNFTHandler(cdc, cliCtx),
	).Methods("POST")
}

type sendNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Denom     string       `json:"denom"`
	ID        string       `json:"id"`
	Recipient string       `json:"recipient"`
}

func transferNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sendNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		recipient, err := sdk.AccAddressFromBech32(req.Recipient)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// create the message
		msg := types.NewMsgSendNFT(cliCtx.GetFromAddress(), recipient, req.Denom, req.ID)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type editNFTMetadataReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
	Name    string       `json:"name"`
}

func editNFTMetadataHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req editNFTMetadataReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgEditNFTMetadata(cliCtx.GetFromAddress(), req.ID, req.Denom, req.TokenURI)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type mintNFTReq struct {
	BaseReq   rest.BaseReq   `json:"base_req"`
	Recipient sdk.AccAddress `json:"recipient"`
	Denom     string         `json:"denom"`
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Hash      string         `json:"hash"`
	Proof     string         `json:"proof"`
}

func mintNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mintNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgMintNFT(cliCtx.GetFromAddress(), req.Recipient, req.ID, req.Denom, req.Name, req.Hash, req.Proof)

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type burnNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
}

func burnNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req burnNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), req.ID, req.Denom)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type buyNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
	Bid     sdk.Coins    `json:"bid"`
}

func buyNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgBuyNFT(cliCtx.GetFromAddress(), req.ID, req.Denom, req.Bid)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type sellNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Denom   string       `json:"denom"`
	ID      string       `json:"id"`
	Ask     sdk.Coins    `json:"ask"`
}

func buyNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sellNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgSellNFT(cliCtx.GetFromAddress(), req.ID, req.Denom, req.Ask)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type challengeNFTReq struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	ContenderDenom string       `json:"contenderdenom"`
	ContenderID    string       `json:"contenderid"`
	DefiantDenom   string       `json:"defiantdenom"`
	DefiantID      string       `json:"defiantid"`
}

func challengeNFTHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req challengeNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgChallengeNFT(cliCtx.GetFromAddress(), req.ID, req.ContenderDenom, req.ContenderID, req.DefiantDenom, req.DefiantID)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type challengeNFTProofReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Denom     string       `json:"denom"`
	ID        string       `json:"id"`
	Recipient string       `json:"recipient"`
}

func challengeNFTProofHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req challengeNFTReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgChallengeNFTProof(cliCtx.GetFromAddress(), req.Recipient, req.Denom, req.ID)
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
