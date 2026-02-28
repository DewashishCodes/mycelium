package cmd

import (
	"fmt"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Sync (rebase) current branch with another branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetBranch := args[0]

		r, _ := git.PlainOpen(".")
		head, _ := r.Head()
		currentBranch := head.Name().Short()

		fmt.Printf("[INFO] Syncing %s with updates from %s...\n", currentBranch, targetBranch)

		// Programmatic rebase is extremely complex in go-git,
		// so for a production-ready tool, we wrap the git CLI
		// to handle complex merge conflicts safely.
		out, err := exec.Command("git", "rebase", targetBranch).CombinedOutput()

		if err != nil {
			fmt.Println("[ERROR] Sync conflict detected!")
			fmt.Println(string(out))
			fmt.Println("[INFO] Use standard git tools to resolve conflicts, or 'git rebase --abort'.")
			return
		}

		fmt.Printf("[SUCCESS] %s is now up-to-date with %s.\n", currentBranch, targetBranch)
	},
}
