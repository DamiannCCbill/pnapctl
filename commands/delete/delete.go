package delete

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/delete/cluster"
	"phoenixnap.com/pnap-cli/commands/delete/privatenetwork"
	"phoenixnap.com/pnap-cli/commands/delete/server"
	"phoenixnap.com/pnap-cli/commands/delete/sshkey"

	serverprivatenetwork "phoenixnap.com/pnap-cli/commands/delete/server/privatenetwork"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource.",
	Long:  `Delete a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DeleteCmd.AddCommand(server.DeleteServerCmd)
	DeleteCmd.AddCommand(serverprivatenetwork.DeleteServerPrivateNetworkCmd)
	DeleteCmd.AddCommand(sshkey.DeleteSshKeyCmd)
	DeleteCmd.AddCommand(cluster.DeleteClusterCmd)
	DeleteCmd.AddCommand(privatenetwork.DeletePrivateNetworkCmd)
}
