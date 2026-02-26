package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

func init() {
	rootCmd.AddCommand(reviewCmd)
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Get an AI recruiter to critique your resume",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := getAPIKey()
		if apiKey == "" {
			fmt.Println("❌ API Key not found. Run: cvvc config --key YOUR_KEY")
			return
		}

		fmt.Println("🤖 AI Recruiter is analyzing your resume...")

		// 1. Read Resume
		resumeData, err := os.ReadFile("resume.json")
		if err != nil {
			fmt.Println("❌ Error: resume.json not found.")
			return
		}

		// 2. Setup Gemini
		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			fmt.Println("❌ Error connecting to AI:", err)
			return
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-2.5-flash")

		// 3. The Prompt (The "Intelligence" part)
		prompt := fmt.Sprintf(`
			Act as a Senior Technical Recruiter at a top-tier tech company (like Google or NVIDIA).
			Review the following resume data provided in JSON format.
			Provide a brutal but constructive critique focused on:
			1. Impact: Are the bullet points quantifying results?
			2. Skills: Are the technologies relevant for an AI/ML role?
			3. Formatting: Is the information clear?
			
			Output the response in a clear, professional terminal-friendly format.
			
			RESUME DATA:
			%s
		`, string(resumeData))

		// 4. Generate
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			fmt.Println("❌ AI Error:", err)
			return
		}

		// 5. Print Result
		fmt.Println("\n--- AI RECRUITER FEEDBACK ---")
		for _, cand := range resp.Candidates {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
		fmt.Println("\n-----------------------------")
	},
}
