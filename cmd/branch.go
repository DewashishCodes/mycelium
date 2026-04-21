package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.AddCommand(branchCreateCmd)
	branchCmd.AddCommand(branchSwitchCmd)
	branchCmd.AddCommand(branchListCmd)
	branchCmd.AddCommand(branchDeleteCmd)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage resume branches (e.g. for different job roles)",
}

var branchListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all resume branches",
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := git.PlainOpen(".")
		branches, _ := r.Branches()

		head, _ := r.Head()
		current := head.Name()

		fmt.Println("🌱 RESUME BRANCHES:")
		branches.ForEach(func(br *plumbing.Reference) error {
			prefix := "  "
			if br.Name() == current {
				prefix = "* "
			}
			fmt.Printf("%s%s\n", prefix, br.Name().Short())
			return nil
		})
	},
}

var branchCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new resume branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := git.PlainOpen(".")
		w, _ := r.Worktree()

		name := args[0]
		branchRef := plumbing.NewBranchReferenceName(name)

		err := w.Checkout(&git.CheckoutOptions{
			Branch: branchRef,
			Create: true,
		})

		if err != nil {
			fmt.Println("Error creating branch:", err)
		} else {
			fmt.Printf("🌱 Branch '%s' created and active.\n", name)
		}
	},
}

var branchSwitchCmd = &cobra.Command{
	Use:   "switch [name]",
	Short: "Switch to a different resume branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := git.PlainOpen(".")
		w, _ := r.Worktree()

		name := args[0]
		err := w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(name),
		})

		if err != nil {
			fmt.Printf("[ERROR] Branch '%s' does not exist.\n", name)
		} else {
			fmt.Printf("🔄 Switched to branch '%s'.\n", name)
		}
	},
}

var branchDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a resume branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := git.PlainOpen(".")
		name := args[0]
		
		head, _ := r.Head()
		if head.Name().Short() == name {
			fmt.Printf("[ERROR] Cannot delete the currently active branch '%s'.\n", name)
			return
		}

		err := r.Storer.RemoveReference(plumbing.NewBranchReferenceName(name))
		if err != nil {
			fmt.Printf("[ERROR] Failed to delete branch '%s': %v\n", name, err)
		} else {
			fmt.Printf("🗑️ Branch '%s' deleted.\n", name)
		}
	},
}

