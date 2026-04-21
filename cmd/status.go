package cmd

import (
	"fmt"
	"mycelium/internal/vcs"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current branch and unsaved changes",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := vcs.Open()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		// 1. Get current branch
		branch, err := r.CurrentBranch()
		if err != nil {
			fmt.Println("📍 Current Branch: (initial branch)")
		} else {
			fmt.Printf("📍 Current Branch: %s\n", branch)
		}

		// 2. Check for changes
		status, err := r.Status()
		if err != nil {
			fmt.Println("[ERROR] Failed to get status:", err)
			return
		}

		if status.IsClean() {
			fmt.Println("[INFO] Mycelium network is healthy and synchronized.")
		} else {
			// Check if it's a real change or just a timestamp change
			if status["resume.json"].Worktree == git.Unmodified {
				fmt.Println("[INFO] Mycelium network is healthy (metadata changes ignored).")
			} else {
				fmt.Println("[WARN] Uncommitted changes detected in the network.")
				fmt.Println("[INFO] Run 'mycelium commit' to protect this version.")
			}
		}
	},
}

