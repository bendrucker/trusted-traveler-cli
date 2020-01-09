package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func printOutput(cmd *cobra.Command, result interface{}) error {
	switch output := viper.GetString("output"); output {
	case "table":
		table := tableprinter.New(cmd.OutOrStdout())
		table.Print(result)
	case "json":
		e := json.NewEncoder(cmd.OutOrStdout())
		e.SetIndent("", "  ")
		return e.Encode(result)
	default:
		return fmt.Errorf("invalid output format: %s", output)
	}

	return nil
}
