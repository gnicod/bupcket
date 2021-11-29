package cmd

import (
	"github.com/gnicod/bupcket/api"
	"github.com/gnicod/bupcket/config"
	"github.com/gnicod/bupcket/storage"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app := api.NewApp(storage.NewS3Provider(), config.GetConfig())
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
