// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/sopstool/execwrap"
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Aliases: []string{"d"},
	Use:     "decrypt [files ...]",
	Short:   "decrypt files",
	Long:    `Decrypt some or all files`,
	RunE:    DecryptCommand,
}

func init() {
	RootCmd.AddCommand(decryptCmd)

	// addCmd.Flags().Bool("nofail", false, "Don't error unless all decrypts fail")
}

// DecryptCommand decrypts files
func DecryptCommand(cmd *cobra.Command, args []string) error {
	filesToDecrypt, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//decrypt all the files
	for _, f := range filesToDecrypt {
		err := execwrap.DecryptFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
