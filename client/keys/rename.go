package keys

import (
	"bufio"
	"errors"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/input"
)

// RenameKeyCommand renames a key from the key store.
func RenameKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename <oldName> <newName>...",
		Short: "Rename the old key name with new key name",
		Long:  `Rename the given key from the Keybase backend.`,
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			buf := bufio.NewReader(cmd.InOrStdin())
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// info, err := clientCtx.Keyring.Key(args[0])
			// if err != nil {
			// 	return err
			// }

			// confirm rename key, unless -y is passed
			if skip, _ := cmd.Flags().GetBool(flagYes); !skip {
				if yes, err := input.GetConfirmation("existing key name will be renamed. Continue?", buf, cmd.ErrOrStderr()); err != nil {
					return err
				} else if !yes {
					return errors.New("aborted")
				}
			}

			if err := clientCtx.Keyring.Rename(args[0], args[1]); err != nil {
				return err
			}

			cmd.Println("Successfully renamed!")
			return nil
		},
	}
	return cmd
}
