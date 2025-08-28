package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var upperAfter bool
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Print a friendly greeting",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := "world"
		nameBefore := name
		nameAfter := name
		if len(args) == 1 {
			name = args[0]
		}
		upperBefore, _ := cmd.Flags().GetBool("upper-before")
		if upperBefore {
			nameBefore = strings.ToUpper(fmt.Sprintf("%s by before", nameBefore))
		}
		if upperAfter {
			nameAfter = strings.ToUpper(fmt.Sprintf("%s by after", nameAfter))
		}
		fmt.Printf("Hello, %s!\n", nameBefore)
		fmt.Printf("Hello, %s!\n", nameAfter)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().Bool("upper-before", false, "Display the name in uppercase before processing")
	greetCmd.Flags().BoolVarP(&upperAfter, "upper-after", "a", false, "Display the name in uppercase after processing")
}
