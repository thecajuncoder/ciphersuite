/*Package cmd contains the CLI sub-commands for cipher suite application
Copyright © 2020 LA Cajun Coder
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a message that has been encoded using a cipher",
	Long: `Decode a message that has been encoded using a cipher. 
The message can be read from a file or, if no file is provided, from STDIN.
Similarly, the decoded message will be written to a file or STDOUT if no output file is specified`,
	Run: runDecodeCommand,
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringVarP(&cipherName, "cipher", "c", "", "The name of the cipher to use")
	decodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The file whose contents will be encoded")
	decodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The file to write the encoded message into")
	decodeCmd.Flags().StringVarP(&cipherKeyStr, "key", "k", "", "The key to use for the specified cipher")
	decodeCmd.Flags().StringVarP(&cipherKeyFile, "keyfile", "x", "", "A file containing the key to use for the specified cipher")

	decodeCmd.MarkFlagRequired("cipher")
}

// runDecodeCommand the "main" method for the "decode" sub-command of the cipher suite CLI
func runDecodeCommand(cmd *cobra.Command, args []string) {

	input, output, cipher, err := checkAndProcessParameters()

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer input.Close()
	defer output.Close()

	_, err = cipher.Decode(input, output)

	if err != nil {
		log.Fatalf("Error decoding message: '%s'\n", err.Error())
	}
}
