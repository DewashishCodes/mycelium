package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/go-git/go-git/v5"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new CVVC repository",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Initialize Git Repo
		_, err := git.PlainInit(".", false)
		if err == git.ErrRepositoryAlreadyExists {
			fmt.Println("⚠️  Repo already exists here!")
		} else if err != nil {
			fmt.Printf("Error initializing git: %s\n", err)
			return
		}

		// 2. The "Jake's Resume" Template
		jakesJson := `{
  "basics": {
    "name": "Jake Ryan",
    "email": "jake@qmail.com",
    "phone": "123-456-7890",
    "linkedin": "linkedin.com/in/jake",
    "github": "github.com/jake",
    "website": "jake-ryan.com"
  },
  "education": [
    {
      "school": "Southwestern University",
      "location": "Georgetown, TX",
      "degree": "Bachelor of Arts in Computer Science",
      "date": "Aug. 2018 – May 2021"
    }
  ],
  "experience": [
    {
      "company": "Undergraduate Research Assistant",
      "role": "Research Assistant",
      "location": "College Station, TX",
      "date": "June 2020 – Present",
      "points": [
        "Developed a REST API using FastAPI and PostgreSQL",
        "Managed database schema using Alembic"
      ]
    }
  ],
  "skills": {
    "languages": ["Java", "Python", "Go", "SQL"],
    "tools": ["Git", "Docker", "VS Code"]
  }
}`

		// 3. Write the file
		err = os.WriteFile("resume.json", []byte(jakesJson), 0644)
		if err != nil {
			fmt.Println("Error creating resume.json:", err)
			return
		}

		// 4. Create .gitignore
		ignoreContent := "*.pdf\ncvvc.exe\n"
		os.WriteFile(".gitignore", []byte(ignoreContent), 0644)

		fmt.Println("✅ CVVC Initialized! Created 'resume.json' with Jake's template.")
	},
}