package cmd

import (
	"fmt"
	"os"

	ttapi "github.com/bendrucker/trusted-traveler-cli/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const baseURL = "https://ttp.cbp.dhs.gov/schedulerapi/"

var rootCmd = &cobra.Command{
	Use:   "trusted-traveler",
	Short: "Interact with the Trusted Traveler Program API",
}

var client = ttapi.NewClient(ttapi.Options{
	URL: baseURL,
})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	f := rootCmd.PersistentFlags()

	f.StringP("output", "o", "table", "Sets the output format (table, json)")

	viper.BindPFlags(f)
}
