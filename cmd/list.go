package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all saved versions (commits)",
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := git.PlainOpen(".")

		logs, err := r.Log(&git.LogOptions{})
		if err != nil {
			fmt.Println("No versions found yet.")
			return
		}

		fmt.Println("🕒 VERSION HISTORY:")
		fmt.Println("-------------------")
		logs.ForEach(func(c *object.Commit) error {
			fmt.Printf("[%s] %s (%s)\n", c.Hash.String()[:7], c.Message, c.Author.When.Format("2006-01-02"))
			return nil
		})
	},
}
