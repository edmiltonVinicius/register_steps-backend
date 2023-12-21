package cli

import (
	"github.com/edmiltonVinicius/register-steps/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server API",
	Long:  "Start the server using .env file localized in root project",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAPI()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
