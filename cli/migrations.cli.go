package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/edmiltonVinicius/register-steps/config"
	"github.com/golang-migrate/migrate"
	"github.com/spf13/cobra"
)

var migrationsCmd = &cobra.Command{
	Use:     "migrations",
	Short:   "Run up, down a migration",
	Long:    "This command can be used to run up or down migration",
	Example: `... migrations -action=up -version=1`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		action, _ := cmd.Flags().GetString("action")
		action = strings.ToLower(action)

		if action != "up" && action != "down" && action != "create" {
			return fmt.Errorf("invalid action. please, use up, down")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(color.InBold("Migrations cli started\n"))

		action, _ := cmd.Flags().GetString("action")
		version, _ := cmd.Flags().GetInt("version")

		run(action, version)
	},
}

func init() {
	rootCmd.AddCommand(migrationsCmd)

	migrationsCmd.Flags().StringP("action", "a", "", "Action to be executed by migrations cli (up, down)")
	migrationsCmd.Flags().IntP("version", "v", 0, "Number Version to be executed by migrations cli")

	migrationsCmd.MarkFlagRequired("action")
}

func logError(err error, act string, version any) {
	log.Fatalf(
		color.InRed(`
			Failed to run migration:
			Action: %s
			Version: %d
			Error: %s
		`),
		act, version, err.Error(),
	)
}

func run(action string, version ...int) {
	config.LoadEnv(false)

	connection, err := config.StartConnectionDB()
	if err != nil {
		log.Fatal("Failed to connect to the Database. \n", err.Error())
	}
	defer connection.Conn.Close()

	if version != nil && version[0] != 0 {
		v := uint(version[0])
		err = connection.Migrate.Migrate(v)
		if err.Error() != migrate.ErrNoChange.Error() {
			logError(err, action, version)
		}
	}

	if strings.ToLower(action) == "up" {
		err = connection.Migrate.Up()
		if err.Error() != migrate.ErrNoChange.Error() {
			logError(err, action, version)
		}
		log.Println(color.InGreen(color.InBold("Migrations UP successfully")))
		return
	}

	if strings.ToLower(action) == "down" {
		err = connection.Migrate.Down()
		if err.Error() != migrate.ErrNoChange.Error() {
			logError(err, action, version)
		}
		log.Println(color.InGreen(color.InBold("Migrations DOWN successfully")))
		return
	}
}
