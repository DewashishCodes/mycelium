package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().BoolP("force", "f", false, "Force restore even if there are unsaved changes")
}

var restoreCmd = &cobra.Command{
	Use:   "restore [hash]",
	Short: "Restore resume.json to a previous version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortHash := args[0]
		force, _ := cmd.Flags().GetBool("force")

		r, _ := git.PlainOpen(".")
		w, _ := r.Worktree()

		// 1. Resolve Short Hash to Full Hash
		// This is the fix! It finds the full 40-char ID from your 7-char input
		fullHash, err := r.ResolveRevision(plumbing.Revision(shortHash))
		if err != nil {
			fmt.Printf("[ERROR] Could not find version [%s]. Check 'mycelium list' for valid hashes.\n", shortHash)
			return
		}

		// 2. Safety Check
		if !force {
			status, _ := w.Status()
			if !status.IsClean() {
				fmt.Println("[WARN] Unsaved changes detected.")
				fmt.Println("[INFO] Use --force to overwrite: mycelium restore " + shortHash + " --force")
				return
			}
		}

		// 3. Restore
		err = w.Checkout(&git.CheckoutOptions{
			Hash:  *fullHash,
			Force: true,
		})

		if err != nil {
			fmt.Printf("[ERROR] Restore failed: %v\n", err)
			return
		}

		fmt.Printf("[SUCCESS] Network restored to version [%s].\n", fullHash.String()[:7])
	},
}
