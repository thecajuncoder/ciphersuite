/*Package cmd contains the CLI sub-commands for cipher suite application
Copyright © 2020 LA Cajun Coder
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thecajuncoder/ciphersuite/cipher"
)

// global parameter used by most or all sub-commands of the CLI app

var cipherKey cipher.Key

// variables used for CLI parameters

var cipherName string
var inputFile string
var outputFile string
var cipherKeyStr string
var cipherKeyFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ciphersuite",
	Short: "A CLI suite of tools for encoding, decoding, and cracking messages using ciphers",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init initialize any global settings (such as a logger) for the entire CLI app
func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}
