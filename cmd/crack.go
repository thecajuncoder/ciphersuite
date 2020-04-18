/*Package cmd contains the CLI sub-commands for cipher suite application
Copyright © 2020 LA Cajun Coder
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// crackCmd represents the crack command
var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crack called")
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
