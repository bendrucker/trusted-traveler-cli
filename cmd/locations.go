package cmd

import (
	ttapi "github.com/bendrucker/trusted-traveler-cli/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// locationsCmd represents the locations command
var locationsCmd = &cobra.Command{
	Use:   "locations",
	Short: "Get locations where Trusted Traveler services are available",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := ttapi.LocationParameters{
			InviteOnly:  ttapi.Bool(viper.GetBool("invite-only")),
			Operational: ttapi.Bool(viper.GetBool("operational")),
		}

		if viper.IsSet("service-name") {
			params.ServiceName = ttapi.String(viper.GetString("service-name"))
		}

		locations, err := client.Locations.List(params)
		if err != nil {
			return err
		}

		return printOutput(cmd, locations)
	},
}

func init() {
	rootCmd.AddCommand(locationsCmd)

	f := locationsCmd.PersistentFlags()

	f.Bool("invite-only", false, "Returns invite-only loctations")
	f.Bool("operational", true, "Returns operational loctations")
	f.String("service-name", "", "Returns loctations that offer the specified service")

	viper.BindPFlags(f)
}
