package cli

import (
	"os"

	"github.com/edmiltonVinicius/register-steps/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   ".",
	Short: "Run the Server API or run migrations Up or Down",
	Long:  "This cli can used to run the server api (default) or management database migrations running up or down.",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAPI()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
