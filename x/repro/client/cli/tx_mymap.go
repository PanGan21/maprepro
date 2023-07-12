package cli

import (
	"maprepro/x/repro/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateMymap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-mymap [index]",
		Short: "Create a new mymap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			mockRuleList := types.RuleList{
				Rule: []types.Rule{
					{
						Rule: map[string]int32{
							"key1": 10,
							"key2": 20,
						},
					},
					{
						Rule: map[string]int32{
							"key3": 30,
							"key4": 40,
						},
					},
					{
						Rule: map[string]int32{
							"key5": 30,
							"key6": 40,
						},
					},
					{
						Rule: map[string]int32{
							"key7": 30,
							"key8": 40,
						},
					},
					{
						Rule: map[string]int32{
							"key9":  30,
							"key10": 40,
						},
					},
				},
			}

			mockCategory := types.Category{
				Rules: map[string]types.RuleList{
					"category1": mockRuleList,
					"category2": mockRuleList,
					"category3": mockRuleList,
					"category4": mockRuleList,
				},
			}

			mockPolicy := types.Policy{
				Policy: map[string]types.Category{
					"policy1": mockCategory,
					"policy2": mockCategory,
					"policy3": mockCategory,
					"policy4": mockCategory,
				},
			}

			msg := types.NewMsgCreateMymap(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				mockPolicy,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMymap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-mymap [index]",
		Short: "Update a mymap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMymap(
				clientCtx.GetFromAddress().String(),
				indexIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMymap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-mymap [index]",
		Short: "Delete a mymap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMymap(
				clientCtx.GetFromAddress().String(),
				indexIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
