package cmd

import (
	"encoding/json"
	"strconv"
	"time"

	ttapi "github.com/bendrucker/trusted-traveler-cli/pkg/api"
	"github.com/olekukonko/tablewriter"
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
			params.LocationID = ttapi.String(viper.GetString("location-id"))
		}

		slots, err := getSlots(params)
		if err != nil {
			return err
		}

		switch viper.GetString("output") {
		case "table":
			table := tablewriter.NewWriter(cmd.OutOrStdout())
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			table.SetAutoMergeCells(true)

			table.SetHeader([]string{
				"Location",
				"Date",
				"Start Time",
			})

			locations, err := client.Locations.List(ttapi.LocationParameters{})
			if err != nil {
				return err
			}

			for _, slot := range slots {
				var location ttapi.Location
				for _, l := range locations {
					if l.ID == slot.LocationID {
						location = l
					}
				}

				zone, err := location.Zone()
				if err != nil {
					return err
				}

				st, err := time.ParseInLocation("2006-01-02T15:04", slot.StartTimestamp, zone)
				if err != nil {
					return err
				}

				table.Append([]string{
					strconv.Itoa(slot.LocationID),
					st.Format("1/2/2006"),
					st.Format("3:04 PM"),
				})
			}

			table.Render()

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
