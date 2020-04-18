/*Package cmd contains the CLI sub-commands for cipher suite application
Copyright © 2020 LA Cajun Coder
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thecajuncoder/ciphersuite/cipher"
	"github.com/thecajuncoder/ciphersuite/util"
)

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

// checkAndProcessParameters performs all of the pre-processing steps for handling encode or decode parameters
// and returns all of the neccessary structs for encoding or decoding with a cipher
// if any error occurrs during pre-processing, it is returned
func checkAndProcessParameters() (io.ReadCloser, io.WriteCloser, cipher.Cipher, error) {

	var cipherKey cipher.Key

	// Parse the cipher key or cipher key file parameters
	if cipherKeyStr == "" && cipherKeyFile == "" {
		return nil, nil, nil, errors.New("Must specify either a Cipher key or key file")
	} else if cipherKeyStr == "" {
		k, err := cipher.ReadCipherKeyFile(cipherKeyFile)

		if err != nil {
			return nil, nil, nil, err
		}
		cipherKey = k
	} else {
		cipherKey = cipher.Key(cipherKeyStr)
	}

	input, err := util.GetInputReader(inputFile)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("Unable to open input file '%s': %s", inputFile, err.Error())
	}

	output, err := util.GetOutputWriter(outputFile)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("Unable to open output file '%s': %s", outputFile, err.Error())
	}

	cipher := util.GetCipherFromName(cipherName)

	if cipher == nil {
		return nil, nil, nil, fmt.Errorf("Cipher '%s' is not a valid Cipher", cipherName)
	}

	cipher.SetKey(cipherKey)
	return input, output, cipher, nil
}
