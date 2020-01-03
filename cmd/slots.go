package cmd

import (
	"encoding/json"
	"time"

	ttapi "github.com/bendrucker/trusted-traveler-cli/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// slotsCmd represents the slots command
var slotsCmd = &cobra.Command{
	Use:   "slots",
	Short: "Lists appointment slots",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := ttapi.SlotParameters{
			OrderBy: ttapi.String(viper.GetString("order-by")),
			Limit:   ttapi.Int(viper.GetInt("limit")),
			Minimum: ttapi.Int(viper.GetInt("minimum")),
		}

		if viper.IsSet("location-id") {
			params.LocationID = ttapi.Int(viper.GetInt("location-id"))
		}

		slots, err := getSlots(params)
		if err != nil {
			return err
		}

		switch viper.GetString("output") {
		case "json":
			e := json.NewEncoder(cmd.OutOrStdout())
			e.SetIndent("", "  ")
			return e.Encode(slots)
		}

		return nil
	},
}

func getSlots(params ttapi.SlotParameters) ([]ttapi.Slot, error) {
	slots, err := client.Slots.List(params)
	if err != nil {
		return nil, err
	}

	if len(slots) == 0 && viper.GetBool("wait") {
		time.Sleep(time.Duration(viper.GetInt("retry-delay")) * time.Second)
		return getSlots(params)
	}

	return slots, err
}

func init() {
	rootCmd.AddCommand(slotsCmd)

	f := slotsCmd.PersistentFlags()

	f.String("order-by", "soonest", "Orders the query for appointment slots")
	f.Int("limit", 10, "The maximum number of slots to return")
	f.Int("minimum", 1, "")
	f.Int("location-id", 0, "Lists appointment slots for the specified location")

	f.Bool("wait", false, "If no slots are returned, retry until slots are returned")
	f.Int("retry-delay", 5, "Number of seconds between retries when --wait is set")

	viper.BindPFlags(f)
}
