package cmd

import (
	"context"
	"fmt"
	"mycelium/internal/resume"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

func init() {
	rootCmd.AddCommand(scoreCmd)
	scoreCmd.Flags().StringP("jd", "j", "", "Path to the Job Description file (.txt)")
}

var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Score your resume against a job description (ATS Simulator)",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := getAPIKey()
		if apiKey == "" {
			fmt.Println("[ERROR] API Key not found. Run: mycelium config --key YOUR_KEY")
			return
		}

		jdPath, _ := cmd.Flags().GetString("jd")
		if jdPath == "" {
			fmt.Println("[ERROR] Please provide a job description file: mycelium score --jd jd.txt")
			return
		}

		jdData, err := os.ReadFile(jdPath)
		if err != nil {
			fmt.Printf("[ERROR] Could not read job description file: %v\n", err)
			return
		}

		resumeData, err := resume.ReadRaw()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		fmt.Println("[AI] Simulating ATS (Applicant Tracking System) scan...")

		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			fmt.Println("[ERROR] AI Connection failed:", err)
			return
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-1.5-flash")

		prompt := fmt.Sprintf(`
			Act as an Applicant Tracking System (ATS) and a Senior Hiring Manager.
			I will provide you with a Resume (JSON) and a Job Description (Text).
			
			Please perform a deep analysis and provide:
			1. **ATS Match Score (0-100%%)**: How well do the keywords and experience align?
			2. **Missing Keywords**: List specific technologies or skills mentioned in the JD but missing from the resume.
			3. **Experience Gap**: Is the seniority level or domain experience a match?
			4. **Recommendations**: Top 3 changes to the resume to increase the score.
			
			RESUME:
			%s
			
			JOB DESCRIPTION:
			%s
		`, string(resumeData), string(jdData))

		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			fmt.Println("[ERROR] AI Generation failed:", err)
			return
		}

		fmt.Println("\n📊 ATS COMPATIBILITY REPORT:")
		fmt.Println("==============================")
		for _, cand := range resp.Candidates {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
		fmt.Println("==============================")
	},
}
