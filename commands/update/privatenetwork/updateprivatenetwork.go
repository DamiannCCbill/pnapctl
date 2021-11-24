package privatenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/networkapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/networkmodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "update private-network"

var Full bool

// UpdatePrivateNetworkCmd is the command for creating a private network.
var UpdatePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network PRIVATE_NETWORK_ID",
	Short:        "Update a private network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update a private network.

Requires a file (yaml or json) containing the information needed to modify the private-network.`,
	Example: `# update a private network as described in privateNetworkUpdate.yaml
pnapctl update private-network 5da891e90ab0c59bd28e34ad --filename ~/privateNetworkUpdate.yaml

# privateNetworkUpdate.yaml
default: true
name: default ssh key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyUpdate, err := networkmodels.CreatePrivateNetworkUpdateFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// update the private network
		response, httpResponse, err := networkapi.Client.PrivateNetworkPut(args[0], *sshKeyUpdate)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
			return printer.PrintPrivateNetworkResponse(response, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	UpdatePrivateNetworkCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	UpdatePrivateNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	UpdatePrivateNetworkCmd.MarkFlagRequired("filename")
}
