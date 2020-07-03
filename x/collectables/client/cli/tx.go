package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/tosch110/collectables/x/collectables/types"
)

// Edit metadata flags
const (
	flagHash  = "hash"
	flagName  = "name"
	flagProof = "proof"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nftTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "NFT transactions subcommands",
		RunE:  client.ValidateCmd,
	}

	nftTxCmd.AddCommand(flags.PostCommands(
		GetCmdSendNFT(cdc),
		GetCmdEditNFTMetadata(cdc),
		GetCmdMintNFT(cdc),
		GetCmdBurnNFT(cdc),
	)...)

	return nftTxCmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdSendNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "Send [sender] [recipient] [denom] [tokenID]",
		Short: "Send a NFT to a recipient",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Send a NFT from a given collection that has a 
			specific id (SHA-256 hex hash) to a specific recipient.
Example:
$ %s tx %s send 
cosmos1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p cosmos1l2rsakp388kuv9k8qzq6lrm9taddae7fpx59wm \
collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			sender, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[2]
			tokenID := args[3]

			msg := types.NewMsgSendNFT(sender, recipient, denom, tokenID)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdEditNFTMetadata is the CLI command for sending an EditMetadata transaction
func GetCmdEditNFTMetadata(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-metadata [denom] [tokenID]",
		Short: "edit the metadata of an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the metadata of an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s edit-metadata collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--name name --from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			tokenID := args[1]
			name := viper.GetString(name)

			msg := types.NewMsgEditNFTMetadata(cliCtx.GetFromAddress(), tokenID, denom, name)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTokenURI, "", "Extra properties available for querying")
	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [denom] [hash] [proof] [name] [recipient]",
		Short: "mint an NFT and set the owner to the recipient",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT from a given collection that has a 
			specific id (SHA-256 hex hash) and set the ownership to a specific address.
Example:
$ %s tx %s mint collectables 03128C68F894E689F009ABD69653297BF111CDC43B9291A5469D8D1C52608C5AEED107A0FE500C3AF7E73F90E3B7CF20 thisismyexampletokeninputforthedemo demo \
cosmos1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p --from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			tokenID := args[1]

			recipient, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			hash := viper.GetString(flagHash)
			name := viper.GetString(flagProof)
			proof := viper.GetString(flagName)

			msg := types.NewMsgMintNFT(cliCtx.GetFromAddress(), recipient, tokenID, denom, hash, proof, name)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTokenURI, "", "URI for supplemental off-chain metadata (should return a JSON object)")

	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burn [denom] [tokenID]",
		Short: "burn an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Burn (i.e permanently delete) an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s burn collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), tokenID, denom)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdBuyNFT is the CLI command for sending a BuyNFT transaction
func GetCmdBuyNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy [denom] [tokenID]",
		Short: "buy an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Buy (i.e make an offer) an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s buy collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgBuyNFT(cliCtx.GetFromAddress(), tokenID, denom)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSellNFT is the CLI command for sending a SellNFT transaction
func GetCmdSellNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sell [denom] [tokenID]",
		Short: "sell an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Sell (i.e make an offer) an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s buy collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgSellNFT(cliCtx.GetFromAddress(), tokenID, denom)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdChallengeNFT is the CLI command for sending a ChallengeNFT transaction
func GetCmdChallengeNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "challenge [contenderdenom] [contendertokenID] [defiantdenom] [defianttokenID]",
		Short: "sell an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Challenge an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s challenge collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			contenderDenom := args[0]
			contenderTokenID := args[1]

			defiantDenom := args[2]
			defiantTokenID := args[3]

			msg := types.NewMsgChallengeNFT(cliCtx.GetFromAddress(), contenderTokenID, contenderDenom, defiantTokenID, defiantDenom)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdChallengeNFTProof is the CLI command for sending a ChallengeNFTProof transaction
func GetCmdChallengeNFTProof(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "challengeproof [contenderdenom] [contendertokenID] [defiantdenom] [defianttokenID]",
		Short: "sell an NFT",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Challenge Proof of an NFT from a given collection that has a 
			specific id (SHA-256 hex hash).
Example:
$ %s tx %s challengeproof collectables d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa \
--from mykey
`,
				version.ClientName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			denom := args[0]
			token := args[1]

			msg := types.NewMsgChallengeNFT(cliCtx.GetFromAddress(), denom, token)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
