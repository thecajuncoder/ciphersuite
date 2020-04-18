/*Package cmd contains the CLI sub-commands for cipher suite application
Copyright © 2020 LA Cajun Coder
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thecajuncoder/ciphersuite/cipher"
	"github.com/thecajuncoder/ciphersuite/util"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Decode a message that has been encoded using a cipher",
	Long: `Decode a message that has been encoded using a cipher. 
The message can be read from a file or, if no file is provided, from STDIN.
Similarly, the encoded message will be written to a file or STDOUT if no output file is specified`,
	Run: runEncodeCommand,
}

// init bind the CLI parameters for the encode command
func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&cipherName, "cipher", "c", "", "The name of the cipher to use")
	encodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The file whose contents will be encoded")
	encodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The file to write the encoded message into")
	encodeCmd.Flags().StringVarP(&cipherKeyStr, "key", "k", "", "The key to use for the specified cipher")
	encodeCmd.Flags().StringVarP(&cipherKeyFile, "keyfile", "x", "", "A file containing the key to use for the specified cipher")

	encodeCmd.MarkFlagRequired("cipher")
}

// runEncodeCommand the "main" method for the "encode" sub-command of the cipher suite CLI
func runEncodeCommand(cmd *cobra.Command, args []string) {

	// Parse the cipher key or cipher key file parameters
	if cipherKeyStr == "" && cipherKeyFile == "" {
		log.Fatalln("Must specify either a Cipher key or key file")
	} else if cipherKeyStr == "" {
		k, err := cipher.ReadCipherKeyFile(cipherKeyFile)

		if err != nil {
			log.Fatalln(err.Error())
		}
		cipherKey = k
	} else {
		cipherKey = cipher.Key(cipherKeyStr)
	}

	input, err := util.GetInputReader(inputFile)

	if err != nil {
		log.Fatalf("Unable to open input file '%s': %s\n", inputFile, err.Error())
	}

	defer input.Close()

	output, err := util.GetOutputWriter(outputFile)

	if err != nil {
		log.Fatalf("Unable to open output file '%s': %s\n", outputFile, err.Error())
	}

	defer output.Close()

	cipher := util.GetCipherFromName(cipherName)

	if cipher == nil {
		log.Fatalf("Cipher '%s' is not a valid Cipher\n", cipherName)
	}

	cipher.SetKey(cipherKey)
	_, err = cipher.Encode(input, output)

	if err != nil {
		log.Fatalf("Error encoding message: '%s'\n", err.Error())
	}
}
