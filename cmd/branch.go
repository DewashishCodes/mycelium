package cmd

import (
	"fmt"
	"mycelium/internal/vcs"

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
		r, err := vcs.Open()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		current, _ := r.CurrentBranch()
		branches, err := r.ListBranches()
		if err != nil {
			fmt.Println("[ERROR] Failed to list branches:", err)
			return
		}

		fmt.Println("🌱 RESUME BRANCHES:")
		for _, name := range branches {
			prefix := "  "
			if name == current {
				prefix = "* "
			}
			fmt.Printf("%s%s\n", prefix, name)
		}
	},
}

var branchCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new resume branch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := vcs.Open()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		name := args[0]
		err = r.CreateBranch(name)
		if err != nil {
			fmt.Printf("[ERROR] Failed to create branch '%s': %v\n", name, err)
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
		r, err := vcs.Open()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		name := args[0]
		err = r.SwitchBranch(name)
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
		r, err := vcs.Open()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		name := args[0]
		err = r.DeleteBranch(name)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		} else {
			fmt.Printf("🗑️ Branch '%s' deleted.\n", name)
		}
	},
}


