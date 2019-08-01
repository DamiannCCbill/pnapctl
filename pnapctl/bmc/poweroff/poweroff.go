package poweroff

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

// ErrorCode to represent the status of the command execution.
// To be only used in testing.
// "OK"     => No errors.
// "ARGS"   => Arg amount error.
// "CLIENT" => Http Client failure error.
// "409"    => 409 response.
// "404"    => 404 response.
// "??"		=> Others
var ErrorCode = "OK"

var P_OffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Powers off a specific server.",
	Long:  "Powers off a specific server.",
	Run: func(cmd *cobra.Command, args []string) {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			ErrorCode = "ARGS"
			return
		}

		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering off server:", err)
			ErrorCode = "CLIENT"
			return
		}

		switch response.StatusCode {
		case 409:
			fmt.Println("Error: Conflict detected. Server is already powered-off.")
			ErrorCode = "409"
		case 404:
			fmt.Println("Error: Server with ID", args[0], "not found.")
			ErrorCode = "404"
		case 500:
			fmt.Println("Error: Internal server error. Please try again later.")
			ErrorCode = "500"
		case 200:
			fmt.Println("Powered off successfully.")
			ErrorCode = "OK"
		default:
			fmt.Println("Status:", response.Status)
			ErrorCode = "??"
		}
	},
}

func init() {
}